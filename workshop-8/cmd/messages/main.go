package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/opentracing/opentracing-go"
	pb "github.com/route256/workshop-8/pkg/messages"
	"github.com/uber/jaeger-client-go/config"
	"google.golang.org/grpc"

	impl "github.com/route256/workshop-8/internal/app/messages"
)

func main() {
	ctx := context.Background()
	flag.Parse()

	var addr string

	flag.StringVar(&addr, "add", ":50051", "Add for messages server")

	if err := run(ctx, addr); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context, addr string) error {
	cfg := config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:            false,
			BufferFlushInterval: 1 * time.Second,
		},
	}
	tracer, closer, err := cfg.New(
		"messages-service",
	)
	if err != nil {
		fmt.Printf("cannot create tracer: %v\n", err)
		os.Exit(1)
	}
	defer closer.Close()

	opentracing.SetGlobalTracer(tracer)

	server := grpc.NewServer()

	pb.RegisterMessagesServer(server, impl.New())

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	log.Printf("service messages listening on %q", addr)
	return server.Serve(lis)
}
