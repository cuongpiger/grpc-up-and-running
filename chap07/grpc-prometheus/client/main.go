package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/cuongpiger/golang/ecommerce"
)

const (
	address = "localhost:50051"
)

func main() {
	// Create a metrics registry.
	reg := prometheus.NewRegistry()
	// Create some standard client metrics.
	grpcMetrics := grpc_prometheus.NewClientMetrics()
	// Register client metrics to registry.
	reg.MustRegister(grpcMetrics)

	// Set up a connection to the server.
	conn, err := grpc.NewClient(address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(grpcMetrics.UnaryClientInterceptor()),
		grpc.WithStreamInterceptor(grpcMetrics.StreamClientInterceptor()),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Create a HTTP server for prometheus.
	httpServer := &http.Server{Handler: promhttp.HandlerFor(reg, promhttp.HandlerOpts{}), Addr: fmt.Sprintf("0.0.0.0:%d", 9094)}

	// Start your http server for prometheus.
	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			log.Fatal("Unable to start a http server.")
		}
	}()

	c := pb.NewProductInfoClient(conn)

	for {
		// Contact the server and print out its response.
		name := "Sumsung S10"
		description := "Samsung Galaxy S10 is the latest smart phone, launched in February 2019"
		price := float32(700.0)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := c.AddProduct(ctx, &pb.Product{Name: name, Description: description, Price: price})
		if err != nil {
			log.Fatalf("Could not add product: %v", err)
		}
		log.Printf("Product ID: %s added successfully", r.Value)

		product, err := c.GetProduct(ctx, &pb.ProductID{Value: r.Value})
		if err != nil {
			log.Fatalf("Could not get product: %v", err)
		}
		log.Printf("Product: ", product.String())
		time.Sleep(3 * time.Second)
	}
}
