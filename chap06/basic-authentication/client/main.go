package main

import (
	"context"
	"encoding/base64"
	"log"
	"os"
	"path/filepath"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/cuongpiger/golang/ecommerce"
)

const (
	address = "localhost:50051"
)

func main() {
	wd, _ := os.Getwd()
	serverCert := filepath.Join(wd, "..", "certs", "server.crt")

	creds, err := credentials.NewClientTLSFromFile(serverCert, "localhost")
	if err != nil {
		log.Fatalf("failed to load credentials: %v", err)
	}
	auth := basicAuth{
		username: "admin",
		password: "admin",
	}
	opts := []grpc.DialOption{
		grpc.WithPerRPCCredentials(auth),
		// transport credentials.
		grpc.WithTransportCredentials(creds),
	}

	// Set up a connection to the server.
	conn, err := grpc.NewClient(address, opts...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewProductInfoClient(conn)

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
}

type basicAuth struct {
	username string
	password string
}

func (b basicAuth) GetRequestMetadata(ctx context.Context, in ...string) (map[string]string, error) {
	auth := b.username + ":" + b.password
	enc := base64.StdEncoding.EncodeToString([]byte(auth))
	return map[string]string{
		"authorization": "Basic " + enc,
	}, nil
}

func (b basicAuth) RequireTransportSecurity() bool {
	return true
}