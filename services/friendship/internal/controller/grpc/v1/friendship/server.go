package friendship

import (
	"buf.build/go/protovalidate"
	"context"
	"fmt"
	"github.com/google/uuid"
	pb "github.com/nazarovnick/chat-platform/services/friendship/pkg/api/grpc/friendship/v1"
	"log"
)

type server struct {
	pb.UnimplementedFriendshipServiceServer
	validator protovalidate.Validator
}

func NewServer() (pb.FriendshipServiceServer, error) {
	srv := &server{}

	validator, err := protovalidate.New(
		protovalidate.WithDisableLazy(),
		protovalidate.WithMessages(
			&pb.SendFriendRequest{},
			&pb.SendFriendResponse{},
			&pb.AcceptFriendRequest{},
			&pb.AcceptFriendResponse{},
			&pb.DeclineFriendRequest{},
			&pb.DeclineFriendResponse{},
			&pb.RemoveFriendRequest{},
			&pb.RemoveFriendResponse{},
			&pb.GetFriendsRequest{},
			&pb.GetFriendsResponse{},
		),
	)

	if err != nil {
		return nil, fmt.Errorf("failed to initialize validator: %w", err)
	}

	srv.validator = validator

	return srv, nil
}

func (s *server) RequestFriend(_ context.Context, req *pb.SendFriendRequest) (*pb.SendFriendResponse, error) {
	log.Printf("Request sent to target user with UUID %v from user with UUID %v", req.GetTargetUserUuid(), req.GetUserUuid())
	if err := s.validator.Validate(req); err != nil {
		return nil, rpcValidationError(err)
	}
	return &pb.SendFriendResponse{
		Status: pb.FriendRequestStatus_FRIEND_REQUEST_STATUS_SUCCESS,
	}, nil
}

func (s *server) AcceptFriend(_ context.Context, req *pb.AcceptFriendRequest) (*pb.AcceptFriendResponse, error) {
	log.Printf("Accept request sent to target user with UUID %v from user with UUID %v", req.GetTargetUserUuid(), req.GetUserUuid())
	if err := s.validator.Validate(req); err != nil {
		return nil, rpcValidationError(err)
	}
	return &pb.AcceptFriendResponse{
		Status: pb.AcceptFriendStatus_ACCEPT_FRIEND_STATUS_SUCCESS,
	}, nil
}

func (s *server) DeclineFriend(_ context.Context, req *pb.DeclineFriendRequest) (*pb.DeclineFriendResponse, error) {
	log.Printf("Decline request sent to target user with UUID %v from user with UUID %v", req.GetTargetUserUuid(), req.GetUserUuid())
	if err := s.validator.Validate(req); err != nil {
		return nil, rpcValidationError(err)
	}
	return &pb.DeclineFriendResponse{
		Status: pb.DeclineFriendStatus_DECLINE_FRIEND_STATUS_SUCCESS,
	}, nil
}

func (s *server) RemoveFriend(_ context.Context, req *pb.RemoveFriendRequest) (*pb.RemoveFriendResponse, error) {
	log.Printf("Remove request sent to target user with UUID %v from user with UUID %v", req.GetTargetUserUuid(), req.GetUserUuid())
	if err := s.validator.Validate(req); err != nil {
		return nil, rpcValidationError(err)
	}
	return &pb.RemoveFriendResponse{
		Status: pb.RemoveFriendStatus_REMOVE_FRIEND_STATUS_SUCCESS,
	}, nil
}

func (s *server) GetFriends(_ context.Context, req *pb.GetFriendsRequest) (*pb.GetFriendsResponse, error) {
	log.Printf("GetFriends request from user with UUID %v", req.GetUserUuid())
	if err := s.validator.Validate(req); err != nil {
		return nil, rpcValidationError(err)
	}

	var uuids []string
	for i := 0; i < 10; i++ {
		uuids = append(uuids, uuid.New().String())
	}

	return &pb.GetFriendsResponse{
		FriendUuids: uuids,
	}, nil
}
