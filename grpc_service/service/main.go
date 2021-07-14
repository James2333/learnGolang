package main

import (
	"context"
	"github.com/gofrs/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"learn101/grpc_service"
	"log"
	"net"
)

const (
	port ="127.0.0.1:20000"
)

type server struct {
	productMap map[string]*grpc_service.Product
}

//添加商品
func (s *server) AddProduct(ctx context.Context, req *grpc_service.Product) (resp *grpc_service.ProductId, err error) {
	resp = &grpc_service.ProductId{}
	out, err := uuid.NewV4()
	if err != nil {
		return resp, status.Errorf(codes.Internal, "err while generate the uuid ", err)
	}

	req.Id = out.String()
	if s.productMap == nil {
		s.productMap = make(map[string]*grpc_service.Product)
	}

	s.productMap[req.Id] = req
	resp.Value = req.Id
	return
}

//获取商品
func (s *server) GetProduct(ctx context.Context, req *grpc_service.ProductId) (resp *grpc_service.Product, err error) {
	if s.productMap == nil {
		s.productMap = make(map[string]*grpc_service.Product)
	}

	resp = s.productMap[req.Value]
	return
}

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Println("net listen err ", err)
		return
	}

	s := grpc.NewServer()
	grpc_service.RegisterProductInfoServer(s, &server{})
	log.Println("start gRPC listen on port " + port)
	if err := s.Serve(listener); err != nil {
		log.Println("failed to serve...", err)
		return
	}
}