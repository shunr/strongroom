package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/shunr/strongroom_server/db"
	pb "github.com/shunr/strongroom_server/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/testdata"
	"log"
	"net"
)

var (
	tls      = flag.Bool("tls", true, "Connection uses TLS if true, else plain TCP")
	certFile = flag.String("cert_file", "", "The TLS cert file")
	keyFile  = flag.String("key_file", "", "The TLS key file")
	port     = flag.Int("port", 443, "The server port")
)

type strongroomServer struct {
}

func (s strongroomServer) CreateAccount(ctx context.Context, request *pb.CreateAccountRequest) (*pb.CreateAccountResponse, error) {

	db.CreateAccount(request.GetUsername(), request.GetDisplayName(), request.GetAuthSalt(), request.GetMukSalt(), request.GetAuthVerifier())
	fmt.Println("CREATED! ", request.GetUsername())
	return &pb.CreateAccountResponse{Username: request.GetUsername()}, nil
}

func newServer() *strongroomServer {
	s := &strongroomServer{}
	return s
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", *port))
	log.Println("Port: ", *port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	if *tls {
		if *certFile == "" {
			*certFile = testdata.Path("server1.pem")
		}
		if *keyFile == "" {
			*keyFile = testdata.Path("server1.key")
		}
		creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
		if err != nil {
			log.Fatalf("Failed to generate credentials %v", err)
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}
	log.Println("GRPC server starting...")
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterStrongroomServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}
