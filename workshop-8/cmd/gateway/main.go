package main

import (
	"context"
	"flag"
	"log"
	"net"

	pb "github.com/route256/workshop-8/pkg/gateway"
	"github.com/route256/workshop-8/pkg/logger"
	messages_pb "github.com/route256/workshop-8/pkg/messages"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	impl "github.com/route256/workshop-8/internal/app/gateway"
)

func main() {
	ctx := context.Background()
	flag.Parse()

	var addr string

	flag.StringVar(&addr, "add", ":50052", "Add for messages server")

	if err := run(ctx, addr); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context, addr string) error {
	zapLogger, err := zap.NewProduction()
	if err != nil {
		return err
	}
	logger.SetGlobal(
		zapLogger.With(zap.String("component", "gateway")),
	)

	conn, err := grpc.Dial(":50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return err
	}

	gw := impl.New(messages_pb.NewMessagesClient(conn))

	server := grpc.NewServer()
	pb.RegisterGatewayServer(server, gw)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	logger.Infof(ctx, "service gateway listening on %q", addr)
	return server.Serve(lis)
}
