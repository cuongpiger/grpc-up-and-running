package main

import (
	"context"
	"log"
	pb "github.com/cuongpiger/golang/ecommerce"
	"time"

	"google.golang.org/grpc/credentials/insecure"

	"google.golang.org/grpc"

	"github.com/cuongpiger/golang/order"
)

const (
	address = "localhost:50051"
)

func main() {
	// Setting up a connection to the server.
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewOrderManagementClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// Get Order
	order.GetOrder(ctx, client)

	// // Add Order
	// log.Println("Adding Order...")
	// order.AddOrder(ctx, client)

	// // Search Order : Server streaming scenario
	// order.SearchOrders(ctx, client)

	// // Update Orders : Client streaming scenario
	// order.UpdateOrders(ctx, client)

	// // Process Order : Bi-di streaming scenario
	// order.ProcessOrders(ctx, client)
}
