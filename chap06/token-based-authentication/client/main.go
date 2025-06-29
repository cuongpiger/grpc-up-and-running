package main

import (
	"context"
	"log"
	"path/filepath"
	"time"
	"os"

	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"

	pb "github.com/cuongpiger/golang/ecommerce"
)

const (
	address  = "localhost:50051"
	hostname = "localhost"
)

func main() {
	// Set up the credentials for the connection.
	perRPC := oauth.NewOauthAccess(fetchToken())
	wd, _ := os.Getwd()

	crtFile := filepath.Join(wd, "..", "certs", "server.crt")
	creds, err := credentials.NewClientTLSFromFile(crtFile, hostname)
	if err != nil {
		log.Fatalf("failed to load credentials: %v", err)
	}
	opts := []grpc.DialOption{
		grpc.WithPerRPCCredentials(perRPC),
		// transport credentials.
		grpc.WithTransportCredentials(creds),
	}

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, opts...)
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
	log.Printf("Product: %v", product.String())
}

func fetchToken() *oauth2.Token {
	return &oauth2.Token{
		AccessToken: "some-secret-token",
	}
}