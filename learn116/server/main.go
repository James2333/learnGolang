package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	user "learn101/learn116"
	"log"
	"net"
)

const (
	port = ":50051"
)

type UserService struct {
	// 实现 User 服务的业务对象
}

func (u *UserService) UserIndex(ctx context.Context, request *user.UserIndexRequest) (*user.UserIndexResponse, error) {
	log.Printf("receive user index request: page %d page_size %d", request.Page, request.PageSize)
	return &user.UserIndexResponse{
		Err:                  0,
		Msg:                  "ok",
		Data:                 []*user.UserEntity{
			{Name: "big_cat", Age: 28},
			{Name: "sqrt_cat", Age: 29},
		},
	},nil
	panic("implement me")
}

func (u *UserService) UserView(ctx context.Context, request *user.UserViewRequest) (*user.UserViewResponse, error) {
	log.Printf("receive user view request: uid %d", request.Uid)
	return &user.UserViewResponse{
		Err:                  0,
		Msg:                  "ok",
		Data:                 &user.UserEntity{
			Name:                 "zbc",
			Age:                  22,
		},
	},nil
	panic("implement me")
}

func (u *UserService) UserPost(ctx context.Context, request *user.UserPostRequest) (*user.UserPostResponse, error) {
	log.Printf("receive user post request: name %s password %s age %d", request.Name, request.Password, request.Age)
	return &user.UserPostResponse{
		Err:                  0,
		Msg:                  "ok",
	},nil
	panic("implement me")
}

func (u *UserService) UserDelete(ctx context.Context, request *user.UserDeleteRequest) (*user.UserDeleteResponse, error) {
	log.Printf("receive user delete request: uid %d", request.Uid)
	return &user.UserDeleteResponse{
		Err:                  0,
		Msg:                  "ok",
	},nil
	panic("implement me")
}

func main() {
	//开启tcp服务器
	listen,err:=net.Listen("tcp",port)
	if err != nil {

	}
	log.Println("start grpc server!")
	//创建RPC服务容器
	grpcServer:=grpc.NewServer()
	//注册userserver服务，传入实现了UserServer接口的东西
	user.RegisterUserServer(grpcServer,&UserService{})

	reflection.Register(grpcServer)
	//grpc服务和tcp服务融合
	err=grpcServer.Serve(listen)

}
