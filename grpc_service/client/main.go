package main

import (
	"context"
	"google.golang.org/grpc"
	"learn101/grpc_service"
	"log"
)

const (
	address = "localhost:20000"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Println("did not connect.", err)
		return
	}
	defer conn.Close()

	client := grpc_service.NewProductInfoClient(conn)
	ctx := context.Background()

	id := AddProduct(ctx, client)
	GetProduct(ctx, client, id)
}

// 添加一个测试的商品
func AddProduct(ctx context.Context, client grpc_service.ProductInfoClient) (id string) {
	aMac := &grpc_service.Product{Name: "Mac Book Pro 2019", Description: "From Apple Inc."}
	productId, err := client.AddProduct(ctx, aMac)
	if err != nil {
		log.Println("add product fail.", err)
		return
	}
	log.Println("add product success, id = ", productId.Value)
	return productId.Value
}

// 获取一个商品
func GetProduct(ctx context.Context, client grpc_service.ProductInfoClient, id string) {
	p, err := client.GetProduct(ctx, &grpc_service.ProductId{Value: id})
	if err != nil {
		log.Println("get product err.", err)
		return
	}
	log.Printf("get prodcut success : %+v\n", p)
}