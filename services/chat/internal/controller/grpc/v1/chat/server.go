package chat

import (
	"context"
	pb "github.com/nazarovnick/chat-platform/services/chat/pkg/api/grpc/chat"
)

type server struct {
	pb.UnimplementedChatServiceServer
}

func NewServer() pb.ChatServiceServer {
	return &server{}
}

func (s *server) SendMessage(ctx context.Context, req *pb.SendMessageRequest) (*pb.SendMessageResponse, error) {
	return &pb.SendMessageResponse{
		MessageId: "testID",
	}, nil
}

func (s *server) GetMessages(ctx context.Context, req *pb.GetMessagesRequest) (*pb.GetMessagesResponse, error) {
	return &pb.GetMessagesResponse{
		Messages: []*pb.Message{
			{
				Id:        "testID",
				ChatId:    "testChatID",
				UserId:    "testUserID",
				Content:   "testContent",
				Timestamp: 3000,
			},
		},
	}, nil
}
