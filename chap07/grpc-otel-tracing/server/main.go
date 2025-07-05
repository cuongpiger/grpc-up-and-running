package main

import (
	"context"
	"log"
	"net"

	"github.com/google/uuid"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	otelcodes "go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/cuongpiger/golang/ecommerce"
)

const (
	port = ":50051"
)

type server struct {
	productMap map[string]*pb.Product
	pb.UnimplementedProductInfoServer
}

func (s *server) AddProduct(ctx context.Context, in *pb.Product) (*pb.ProductID, error) {
	tr := otel.Tracer("ecommerce")
	ctx, span := tr.Start(ctx, "AddProduct")
	defer span.End()

	out, err := uuid.NewUUID()
	if err != nil {
		span.SetStatus(otelcodes.Error, err.Error())
		return nil, err
	}

	in.Id = out.String()
	if s.productMap == nil {
		s.productMap = make(map[string]*pb.Product)
	}
	s.productMap[in.Id] = in
	span.SetStatus(otelcodes.Ok, "Product added successfully")
	return &pb.ProductID{Value: in.Id}, nil
}

func (s *server) GetProduct(ctx context.Context, in *pb.ProductID) (*pb.Product, error) {
	tr := otel.Tracer("ecommerce")
	ctx, span := tr.Start(ctx, "GetProduct")
	defer span.End()

	value, exists := s.productMap[in.Value]
	if exists {
		span.SetStatus(otelcodes.Ok, "Product retrieved successfully")
		return value, status.New(codes.OK, "").Err()
	}

	errMsg := "Product does not exist: " + in.Value
	span.SetStatus(otelcodes.Error, errMsg)
	return nil, status.Errorf(codes.NotFound, errMsg)
}

func main() {
	shutdown := initTracing()
	defer shutdown()

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer(
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
	)

	pb.RegisterProductInfoServer(grpcServer, &server{})
	log.Printf("🚀 gRPC Server running on port %v\n", port)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func initTracing() func() {
	exp, err := otlptracegrpc.New(context.Background(), otlptracegrpc.WithEndpointURL("http://localhost:14268/api/traces"))
	if err != nil {
		log.Fatalf("failed to initialize OTLP exporter: %v", err)
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName("product_info"),
		)),
	)

	otel.SetTracerProvider(tp)

	return func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Fatalf("error shutting down tracing: %v", err)
		}
	}
}
