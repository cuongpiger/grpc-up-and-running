package main

import (
	"context"
	"log"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/cuongpiger/golang/ecommerce"
)

const (
	address = "localhost:50051"
)

func setupTracing() func(context.Context) error {
	exp, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	if err != nil {
		log.Fatalf("failed to initialize exporter: %v", err)
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName("grpc-ecommerce-client"),
		)),
	)

	otel.SetTracerProvider(tp)

	return tp.Shutdown
}

func main() {
	ctx := context.Background()

	// Setup tracing
	shutdown := setupTracing()
	defer func() {
		if err := shutdown(ctx); err != nil {
			log.Fatalf("error shutting down tracing: %v", err)
		}
	}()

	// Create gRPC connection with OpenTelemetry interceptors
	conn, err := grpc.NewClient(address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithStatsHandler(otelgrpc.NewClientHandler()),
	)
	if err != nil {
		log.Fatalf("Can't connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewProductInfoClient(conn)

	for {
		name := "Samsung S10"
		description := "Samsung Galaxy S10 is the latest smartphone, launched in February 2019"
		price := float32(700.0)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		// Add product
		r, err := c.AddProduct(ctx, &pb.Product{Name: name, Description: description, Price: price})
		if err != nil {
			log.Fatalf("Could not add product: %v", err)
		}
		log.Printf("✅ Product ID: %s added successfully", r.Value)

		// Get product
		product, err := c.GetProduct(ctx, &pb.ProductID{Value: r.Value})
		if err != nil {
			log.Fatalf("Could not get product: %v", err)
		}
		log.Printf("✅ Product retrieved: %v", product)

		time.Sleep(3 * time.Second)
	}
}
