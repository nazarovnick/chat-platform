package chat

import (
	"context"
	"fmt"
	"github.com/bufbuild/protovalidate-go"
	pb "github.com/nazarovnick/chat-platform/services/chat/pkg/api/grpc/chat"
	"log"
)

type server struct {
	pb.UnimplementedChatServiceServer
	validator protovalidate.Validator
}

func NewServer() (pb.ChatServiceServer, error) {
	srv := &server{}

	validator, err := protovalidate.New(
		protovalidate.WithDisableLazy(),
		protovalidate.WithMessages(
			&pb.SendMessageRequest{},
			&pb.SendMessageResponse{},
			&pb.GetMessagesRequest{},
			&pb.GetMessagesResponse{},
		),
	)

	if err != nil {
		return nil, fmt.Errorf("failed to initialize validator: %w", err)
	}

	srv.validator = validator

	return srv, nil
}

func (s *server) SendMessage(_ context.Context, req *pb.SendMessageRequest) (*pb.SendMessageResponse, error) {
	log.Printf("Message sent to chat with ID %v from user with ID %v", req.GetChatId(), req.GetUserId())
	if err := s.validator.Validate(req); err != nil {
		return nil, rpcValidationError(err)
	}
	return &pb.SendMessageResponse{
		MessageId: 777,
	}, nil
}

func (s *server) GetMessages(_ context.Context, req *pb.GetMessagesRequest) (*pb.GetMessagesResponse, error) {
	log.Printf("Get messages from chat ID %v", req.GetChatId())
	if err := s.validator.Validate(req); err != nil {
		return nil, rpcValidationError(err)
	}

	return &pb.GetMessagesResponse{
		Messages: []*pb.Message{
			{
				Id:        777,
				ChatId:    888,
				UserId:    999,
				Content:   "testContent",
				Timestamp: 3000,
			},
		},
	}, nil
}
