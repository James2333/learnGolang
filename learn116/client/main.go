package main

import (
	"context"
	"google.golang.org/grpc"
	user "learn101/learn116"
	"log"
	"time"
)

const (
	address = "localhost:50051"
)

func main() {
	//创建一个grpc连接
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {

	}
	defer conn.Close()
	//创建RPC客户端容器
	userClient := user.NewUserClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	userIndexReponse, err := userClient.UserIndex(ctx, &user.UserIndexRequest{Page: 1, PageSize: 12,
	})
	if err != nil {

	}
	if userIndexReponse.Err == 0 {
		log.Printf("user index success: %s", userIndexReponse.Msg)
		// 包含 UserEntity 的数组列表
		userEntityList := userIndexReponse.Data
		for _, row := range userEntityList {
			log.Println(row.Name, row.Age)
		}
	} else {
		log.Printf("user index error: %d", userIndexReponse.Err)
	}

	// UserView 请求
	userViewResponse, err := userClient.UserView(ctx, &user.UserViewRequest{Uid: 1})
	if err != nil {
		log.Printf("user view could not greet: %v", err)
	}

	if 0 == userViewResponse.Err {
		log.Printf("user view success: %s", userViewResponse.Msg)
		userEntity := userViewResponse.Data
		log.Println(userEntity.Name, userEntity.Age)
	} else {
		log.Printf("user view error: %d", userViewResponse.Err)
	}

	// UserPost 请求
	userPostReponse, err := userClient.UserPost(ctx, &user.UserPostRequest{Name: "big_cat", Password: "123456", Age: 29})
	if err != nil {
		log.Printf("user post could not greet: %v", err)
	}

	if 0 == userPostReponse.Err {
		log.Printf("user post success: %s", userPostReponse.Msg)
	} else {
		log.Printf("user post error: %d", userPostReponse.Err)
	}

	// UserDelete 请求
	userDeleteReponse, err := userClient.UserDelete(ctx, &user.UserDeleteRequest{Uid: 1})
	if err != nil {
		log.Printf("user delete could not greet: %v", err)
	}

	if 0 == userDeleteReponse.Err {
		log.Printf("user delete success: %s", userDeleteReponse.Msg)
	} else {
		log.Printf("user delete error: %d", userDeleteReponse.Err)
	}

}
