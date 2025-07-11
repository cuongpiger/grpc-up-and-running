// Go to ${grpc-up-and-running}/samples/ch02/productinfo
// Optional: Execute protoc -I proto proto/product_info.proto --go_out=plugins=grpc:go/product_info
// Execute go get -v github.com/grpc-up-and-running/samples/ch02/productinfo/go/product_info
// Execute go run go/server/main.go

package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/google/uuid"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"

	pb "github.com/cuongpiger/golang/ecommerce"
)

const (
	port = ":50051"
)

// server is used to implement ecommerce/product_info.
type server struct {
	productMap map[string]*pb.Product

	pb.UnimplementedProductInfoServer
}

// AddProduct implements ecommerce.AddProduct
func (s *server) AddProduct(ctx context.Context, in *pb.Product) (*pb.ProductID, error) {
	customizedCounterMetric.WithLabelValues(in.Name).Inc()
	out, err := uuid.NewUUID()
	if err != nil {
		log.Fatal(err)
	}
	in.Id = out.String()
	if s.productMap == nil {
		s.productMap = make(map[string]*pb.Product)
	}
	s.productMap[in.Id] = in
	return &pb.ProductID{Value: in.Id}, nil
}

// GetProduct implements ecommerce.GetProduct
func (s *server) GetProduct(ctx context.Context, in *pb.ProductID) (*pb.Product, error) {
	value, exists := s.productMap[in.Value]
	if exists {
		return value, nil
	}
	return nil, errors.New("Product does not exist for the ID" + in.Value)
}

var (
	// Create a metrics registry.
	reg = prometheus.NewRegistry()

	// Create some standard server metrics.
	grpcMetrics = grpc_prometheus.NewServerMetrics()

	customizedCounterMetric = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "product_mgt_server_handle_count",
		Help: "Total number of RPCs handled on the server.",
	}, []string{"name"})
)

func init() {
	// Register standard server metrics and customized metrics to registry.
	reg.MustRegister(grpcMetrics, customizedCounterMetric)
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create a HTTP server for prometheus.
	httpServer := &http.Server{Handler: promhttp.HandlerFor(reg, promhttp.HandlerOpts{}), Addr: fmt.Sprintf("0.0.0.0:%d", 9092)}

	// Create a gRPC Server with gRPC interceptor.
	grpcServer := grpc.NewServer(
		grpc.StreamInterceptor(grpcMetrics.StreamServerInterceptor()),
		grpc.UnaryInterceptor(grpcMetrics.UnaryServerInterceptor()),
	)

	pb.RegisterProductInfoServer(grpcServer, &server{})
	// Initialize all metrics.
	grpcMetrics.InitializeMetrics(grpcServer)

	// Start your http server for prometheus.
	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			log.Fatal("Unable to start a http server.")
		}
	}()

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
