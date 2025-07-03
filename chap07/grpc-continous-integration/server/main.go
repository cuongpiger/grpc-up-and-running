// Go to ${grpc-up-and-running}/samples/ch02/productinfo
// Optional: Execute protoc -I proto-gen proto-gen/product_info.proto-gen --go_out=plugins=grpc:go/product_info
// Execute go get -v github.com/grpc-up-and-running/samples/ch02/productinfo/go/product_info
// Execute go run go/server/main.go

package main

import (
	"context"
	"errors"
	"log"
	"net"
	"sync"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/cuongpiger/golang/ecommerce"
)

const (
	port = ":50051"
)

type productMap struct {
	sync.RWMutex
	products map[string]*pb.Product
}

func newProductMap() *productMap {
	return &productMap{
		RWMutex:  sync.RWMutex{},
		products: make(map[string]*pb.Product),
	}
}

func (s *productMap) Get(id string) (*pb.Product, bool) {
	s.RLock()
	defer s.RUnlock()
	product, exists := s.products[id]
	return product, exists
}

func (s *productMap) Set(id string, product *pb.Product) {
	s.Lock()
	defer s.Unlock()

	s.products[id] = product
	log.Printf("Product added - ID: %s, Name: %s", id, product.Name)
}

// server is used to implement ecommerce/product_info.
type server struct {
	productMap *productMap

	pb.UnimplementedProductInfoServer
}

// AddProduct implements ecommerce.AddProduct
func (s *server) AddProduct(ctx context.Context, in *pb.Product) (*pb.ProductID, error) {
	out, err := uuid.NewUUID()
	if err != nil {
		log.Fatal(err)
	}
	in.Id = out.String()
	if s.productMap == nil {
		s.productMap = newProductMap()
	}
	s.productMap.Set(in.Id, in)
	log.Printf("New product added - ID : %s, Name : %s", in.Id, in.Name)
	return &pb.ProductID{Value: in.Id}, nil
}

// GetProduct implements ecommerce.GetProduct
func (s *server) GetProduct(ctx context.Context, in *pb.ProductID) (*pb.Product, error) {
	value, exists := s.productMap.Get(in.Value)
	if exists {
		log.Printf("New product retrieved - ID : %s", in)
		return value, nil
	}

	return nil, errors.New("Product does not exist for the ID" + in.Value)
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterProductInfoServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
