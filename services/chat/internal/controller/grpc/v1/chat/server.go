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
	log.Printf("Message sent to chat with ID %v from user with UUID %v", req.GetChatId(), req.GetUserUuid())
	if err := s.validator.Validate(req); err != nil {
		return nil, rpcValidationError(err)
	}
	return &pb.SendMessageResponse{
		MessageUuid: "67f3c9a0-b0a2-4ae8-bc52-ff3f4b02e8a1",
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
				Uuid:      "67f3c9a0-b0a2-4ae8-bc52-ff3f4b02e9a3",
				ChatId:    888,
				UserUuid:  "d290f1ee-6c54-4b01-90e6-d701748f0951",
				Content:   "testContent1",
				Timestamp: 3000,
			},
			{
				Uuid:      "17f3c9a0-b0a2-4ae8-bc52-ff3f4b02e8a1",
				ChatId:    999,
				UserUuid:  "1290f1ee-6c54-4b01-90e6-d701748f0851",
				Content:   "testContent2",
				Timestamp: 3000,
			},
		},
	}, nil
}
