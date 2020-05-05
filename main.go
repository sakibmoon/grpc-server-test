package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"example.com/sakib/server/api"
)

type server struct{}

func (s *server) HelloWorld(ctx context.Context, r *api.HelloRequest) (*api.HelloResponse, error) {
	fmt.Println("It came here")
	result := &api.HelloResponse{}
	result.Ack = "This is hello grpc"

	return result, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi, there, I love %s!", r.URL.Path[1:])
}

func main() {
	// http.HandleFunc("/", handler)
	// log.Fatal(http.ListenAndServe(":8080", nil))

	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("Failed to listen:  %v", err)
	}

	s := grpc.NewServer()
	api.RegisterHelloServiceServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
