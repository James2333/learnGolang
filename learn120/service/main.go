package main

import (
	"google.golang.org/grpc"
	stream "learn101/learn120"
	"log"
	"net"
)

const (
	port = "127.0.0.1:20000"
)

type server struct{}

func (s *server) List(r *stream.StreamRequest, listServer stream.StreamService_ListServer) error {
	for n := 0; n <= 6; n++ {
		err := listServer.Send(&stream.StreamResponse{
			Pt: &stream.StreamPoint{
				Name:  r.Pt.Name,
				Value: r.Pt.Value + int32(n),
			},
		})
		if err != nil {
			return err
		}
	}
	return nil
	panic("implement me")

}

func (s *server) Record(recordServer stream.StreamService_RecordServer) error {
	panic("implement me")
}

func (s *server) Route(routeServer stream.StreamService_RouteServer) error {
	panic("implement me")
}

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Println("net listen err ", err)
		return
	}

	s := grpc.NewServer()
	stream.RegisterStreamServiceServer(s, &server{})
	log.Println("start gRPC listen on port " + port)
	if err = s.Serve(listener); err != nil {
		log.Println("failed to serve...", err)
		return
	}
}
