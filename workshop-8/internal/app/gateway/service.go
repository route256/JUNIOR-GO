package gateway

import (
	"context"
	"io"

	pb "github.com/route256/workshop-8/pkg/gateway"
	"github.com/route256/workshop-8/pkg/logger"
	pb_messages "github.com/route256/workshop-8/pkg/messages"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Implementation struct {
	pb.UnimplementedGatewayServer

	messages pb_messages.MessagesClient
}

func New(messages pb_messages.MessagesClient) *Implementation {
	return &Implementation{
		messages: messages,
	}
}

func (i *Implementation) GetMessagesSummary(ctx context.Context, empt *emptypb.Empty) (*pb.MessagesSummary, error) {
	summary, err := i.messages.GetMessagesSummary(ctx, empt)
	if err != nil {
		return nil, err
	}

	return &pb.MessagesSummary{
		Count: summary.GetCount(),
	}, nil
}

func (i *Implementation) PullMessages(ctx context.Context, _ *emptypb.Empty) (*pb.PullMessagesResponse, error) {
	stream, err := i.messages.PullMessages(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}

	resp := pb.PullMessagesResponse{}
	for {
		msg, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		resp.Messages = append(resp.Messages, &pb.Message{
			Ts:     msg.Ts,
			Text:   msg.Text,
			Author: msg.Author,
		})
	}

	return &resp, nil
}

func (i *Implementation) PushMessages(ctx context.Context, req *pb.PushMessagesRequest) (*pb.PushMessagesRequestResponse, error) {
	stream, err := i.messages.PushMessages(ctx)
	if err != nil {
		return nil, err
	}

	for _, message := range req.GetMessages() {
		stream.Send(&pb_messages.Message{
			Ts:     message.Ts,
			Text:   message.Text,
			Author: message.Author,
		})
	}

	summary, err := stream.CloseAndRecv()
	if err != nil {
		return nil, err
	}

	return &pb.PushMessagesRequestResponse{
		Summary: &pb.MessagesSummary{
			Count: summary.GetCount(),
		},
	}, nil
}

func (i *Implementation) ExchangeMessages(ctx context.Context, req *pb.ExchangeMessagesRequest) (*pb.ExchangeMessagesResponse, error) {
	stream, err := i.messages.ExchangeMessages(ctx)
	if err != nil {
		return nil, err
	}

	l := logger.FromContext(ctx)
	ctx = logger.ToContext(ctx, l.With(zap.String("method", "exchange")))

	logger.Infof(ctx, "this message has 'method' attr")

	var resp pb.ExchangeMessagesResponse
	var respErr error

	waitCh := make(chan struct{})
	go func() {
		for {
			message, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					close(waitCh)
					return
				}
				respErr = err
				return
			}

			resp.Messages = append(resp.Messages,
				&pb.Message{
					Ts:     message.Ts,
					Text:   message.Text,
					Author: message.Author,
				},
			)
		}
	}()

	for _, message := range req.GetMessages() {
		stream.Send(&pb_messages.Message{
			Ts:     message.Ts,
			Text:   message.Text,
			Author: message.Author,
		})
	}

	if err := stream.CloseSend(); err != nil {
		return nil, err
	}

	<-waitCh
	return &resp, respErr
}
