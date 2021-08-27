package main

import (
	"context"
	"google.golang.org/grpc"
	"io"
	stream "learn101/learn120"
	//pb "learn101/grpcStream"
	"log"
)

const (
	address = "127.0.0.1:20000"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Println("did not connect.", err)
		return
	}
	//log.Println(conn.)
	defer conn.Close()

	client := stream.NewStreamServiceClient(conn)
	err = printLists(client, &stream.StreamRequest{Pt: &stream.StreamPoint{Name: "gRPC Stream Client: List", Value: 2018}})
	if err != nil {
		log.Fatalf("printLists.err: %v", err)
	}

	//err = printRecord(client, &stream.StreamRequest{Pt: &stream.StreamPoint{Name: "gRPC Stream Client: Record", Value: 2018}})
	//if err != nil {
	//	log.Fatalf("printRecord.err: %v", err)
	//}
	//
	//err = printRoute(client, &stream.StreamRequest{Pt: &stream.StreamPoint{Name: "gRPC Stream Client: Route", Value: 2018}})
	//if err != nil {
	//	log.Fatalf("printRoute.err: %v", err)
	//}
	//ctx := context.Background()

	//id := AddProduct(ctx, client)
	//GetProduct(ctx, client, id)
	//time.Sleep(time.Second*2)
}

func printLists(client stream.StreamServiceClient, r *stream.StreamRequest) error {
	log.Println("Recv msg!")
	s,err:=client.List(context.Background(),r)
	if err != nil {
		return err
	}
	for {
		resp,err:=s.Recv()
		if err != io.EOF {
			break
		}
		if err != nil {
			return err
		}
		//log.Println("im sb")
		log.Printf("resp: pj.name: %s, pt.value: %d", resp.Pt.Name, resp.Pt.Value)
	}
	return nil
}

func printRecord(client stream.StreamServiceClient, r *stream.StreamRequest) error {
	return nil
}

func printRoute(client stream.StreamServiceClient, r *stream.StreamRequest) error {
	return nil
}