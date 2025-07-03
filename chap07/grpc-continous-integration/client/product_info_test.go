package main

import (
	"context"
	"fmt"
	"testing"
	"time"

	"go.uber.org/mock/gomock"
	"google.golang.org/protobuf/proto"

	pb "github.com/cuongpiger/golang/ecommerce"
	mocks "github.com/cuongpiger/golang/mocks"
)

// rpcMsg implements the gomock.Matcher interface
type rpcMsg struct {
	msg proto.Message
}

func (r *rpcMsg) Matches(msg interface{}) bool {
	m, ok := msg.(proto.Message)
	if !ok {
		return false
	}
	return proto.Equal(m, r.msg)
}

func (r *rpcMsg) String() string {
	return fmt.Sprintf("is %s", r.msg)
}

func TestAddProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProdInfoClient := mocks.NewMockProductInfoClient(ctrl)

	name := "Sumsung S10"
	description := "Samsung Galaxy S10 is the latest smart phone, launched in February 2019"
	price := float32(700.0)
	req := &pb.Product{Name: name, Description: description, Price: price}

	mockProdInfoClient.
		EXPECT().AddProduct(gomock.Any(), &rpcMsg{msg: req}).
		Return(&pb.ProductID{Value: "Product:" + name}, nil)

	testAddProduct(t, mockProdInfoClient)

}

func testAddProduct(t *testing.T, client pb.ProductInfoClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	name := "Sumsung S10"
	description := "Samsung Galaxy S10 is the latest smart phone, launched in February 2019"
	price := float32(700.0)

	r, err := client.AddProduct(ctx, &pb.Product{Name: name, Description: description, Price: price})

	if err != nil || r.GetValue() != "Product:Sumsung S10" {
		t.Errorf("mocking failed")
	}
	t.Log("Reply : ", r.GetValue())

}
