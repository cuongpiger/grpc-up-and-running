package main

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"

	"github.com/google/uuid"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/contrib/zpages"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	stdout "go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	"google.golang.org/grpc"

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
	out, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	in.Id = out.String()
	if s.productMap == nil {
		s.productMap = make(map[string]*pb.Product)
	}
	s.productMap[in.Id] = in
	return &pb.ProductID{Value: in.Id}, nil
}

func (s *server) GetProduct(ctx context.Context, in *pb.ProductID) (*pb.Product, error) {
	value, exists := s.productMap[in.Value]
	if exists {
		return value, nil
	}
	return nil, errors.New("Product does not exist for the ID: " + in.Value)
}

// setupTracing initializes tracing with stdout exporter
func setupTracing() func(context.Context) error {
	exp, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	if err != nil {
		log.Fatalf("failed to initialize stdout exporter: %v", err)
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName("grpc-ecommerce-service"),
		)),
	)

	otel.SetTracerProvider(tp)

	return tp.Shutdown
}

func initTracer(zsp *zpages.SpanProcessor) (*sdktrace.TracerProvider, error) {
	// Create stdout exporter to be able to retrieve
	// the collected spans.
	exporter, err := stdout.New()
	if err != nil {
		return nil, err
	}

	// For the demonstration, use sdktrace.AlwaysSample sampler to sample all traces.
	// In a production application, use sdktrace.ProbabilitySampler with a desired probability.
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resource.NewWithAttributes(semconv.SchemaURL, semconv.ServiceNameKey.String("ExampleService"))),
		sdktrace.WithSpanProcessor(zsp),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	return tp, err
}

func main() {
	ctx := context.Background()

	// Setup OpenTelemetry tracing
	shutdown := setupTracing()
	defer func() {
		if err := shutdown(ctx); err != nil {
			log.Fatalf("error shutting down tracing: %v", err)
		}
	}()

	// Set up zPages for tracing visualization/debugging
	zsp := zpages.NewSpanProcessor()
	tp, err := initTracer(zsp)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		http.Handle("/debug/zpages/tracez", zpages.NewTracezHandler(zsp))
		log.Println("View tracez at http://localhost:7777/debug/zpages/tracez, public API at http://localhost:7777/hello")
		err = http.ListenAndServe("localhost:7777", nil)
		if err != nil {
			log.Fatal(err)
		}
	}()

	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()

	// Create a gRPC Server with OpenTelemetry interceptors
	grpcServer := grpc.NewServer(
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
	)

	pb.RegisterProductInfoServer(grpcServer, &server{})

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Printf("🚀 gRPC Server running on port %v\n", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
