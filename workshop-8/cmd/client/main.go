package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/route256/workshop-8/pkg/gateway"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

func main() {
	ctx := context.Background()
	flag.Parse()

	var addr string

	flag.StringVar(&addr, "add", ":50052", "Add for messages server")

	if err := run(ctx, addr, flag.Arg(0)); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context, addr string, cmd string) error {
	conn, err := grpc.Dial(addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}

	client := pb.NewGatewayClient(conn)

	switch cmd {
	case "summary":
		return summary(ctx, client)
	case "pull":
		return pull(ctx, client)
	case "push":
		return push(ctx, client)
	case "exchange":
		return exchange(ctx, client)
	default:
		log.Fatalf("unknown cmd %s (use: summary, pull, push, exchange)", cmd)
	}

	return nil
}

func summary(ctx context.Context, client pb.GatewayClient) error {
	summary, err := client.GetMessagesSummary(ctx, &emptypb.Empty{})
	if err != nil {
		return err
	}

	log.Printf("count %d", summary.GetCount())
	return nil
}

func pull(ctx context.Context, client pb.GatewayClient) error {
	resp, err := client.PullMessages(ctx, &emptypb.Empty{})
	if err != nil {
		return err
	}

	for _, message := range resp.GetMessages() {
		log.Printf("message %q from %s", message.Text, message.Author)
	}

	return nil
}

func push(ctx context.Context, client pb.GatewayClient) error {
	resp, err := client.PushMessages(ctx, &pb.PushMessagesRequest{
		Messages: data,
	})
	if err != nil {
		return err
	}

	log.Printf("pushed %d messages", resp.GetSummary().GetCount())
	return nil
}

func exchange(ctx context.Context, client pb.GatewayClient) error {
	resp, err := client.ExchangeMessages(ctx, &pb.ExchangeMessagesRequest{
		Messages: data,
	})
	if err != nil {
		return err
	}

	for _, message := range resp.GetMessages() {
		log.Printf("message %q from %s", message.Text, message.Author)
	}

	return nil
}

var data = []*pb.Message{
	{
		Ts:     time.Now().Format(time.RFC3339),
		Text:   "Message 1",
		Author: "client",
	},
	{
		Ts:     time.Now().Format(time.RFC3339),
		Text:   "Message 2",
		Author: "client",
	},
	{
		Ts:     time.Now().Format(time.RFC3339),
		Text:   "Message 3",
		Author: "client",
	},
	{
		Ts:     time.Now().Format(time.RFC3339),
		Text:   "Message 4",
		Author: "client",
	},
	{
		Ts:     time.Now().Format(time.RFC3339),
		Text:   "Message 5",
		Author: "client",
	},
}
