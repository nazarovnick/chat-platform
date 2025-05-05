package user

import (
	"buf.build/go/protovalidate"
	"context"
	"fmt"
	pb "github.com/nazarovnick/chat-platform/services/user/pkg/api/grpc/user/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"time"
)

type server struct {
	pb.UnimplementedUserServiceServer
	validator protovalidate.Validator
}

func NewServer() (pb.UserServiceServer, error) {
	srv := &server{}

	validator, err := protovalidate.New(
		protovalidate.WithDisableLazy(),
		protovalidate.WithMessages(
			&pb.GetUserRequest{},
			&pb.GetUserResponse{},
			&pb.GetUserByNicknameRequest{},
			&pb.GetUserByNicknameResponse{},
			&pb.UpdateUserRequest{},
			&pb.UpdateUserResponse{},
			&pb.SearchUsersByNicknameRequest{},
			&pb.SearchUsersByNicknameResponse{},
		),
	)

	if err != nil {
		return nil, fmt.Errorf("failed to initialize validator: %w", err)
	}

	srv.validator = validator

	return srv, nil
}

func (s *server) GetUser(_ context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	log.Printf("GetUser request sent with UUID %v", req.GetUserUuid())
	if err := s.validator.Validate(req); err != nil {
		return nil, rpcValidationError(err)
	}
	return &pb.GetUserResponse{
		UserUuid:    "123e4567-e89b-12d3-a456-426614174000",
		Nickname:    "vasyan888",
		Age:         30,
		BirthDate:   timestamppb.New(time.Date(1993, 5, 21, 0, 0, 0, 0, time.UTC)),
		AvatarUrl:   "https://example.com/avatar.png",
		CreatedAt:   timestamppb.New(time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC)),
		SoftDeleted: false,
	}, nil
}

func (s *server) GetUserByNickname(_ context.Context, req *pb.GetUserByNicknameRequest) (*pb.GetUserByNicknameResponse, error) {
	log.Printf("GetUserByNickname request sent with nickname %v", req.GetNickname())
	if err := s.validator.Validate(req); err != nil {
		return nil, rpcValidationError(err)
	}
	return &pb.GetUserByNicknameResponse{
		UserUuid:    "123e4567-e89b-12d3-a456-426614174000",
		Nickname:    "vasyan888",
		Age:         30,
		BirthDate:   timestamppb.New(time.Date(1993, 5, 21, 0, 0, 0, 0, time.UTC)),
		AvatarUrl:   "https://example.com/avatar.png",
		CreatedAt:   timestamppb.New(time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC)),
		SoftDeleted: false,
	}, nil
}

func (s *server) UpdateUser(_ context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	log.Printf("UpdateUser request sent for user with UUID %v", req.GetUserUuid())
	if err := s.validator.Validate(req); err != nil {
		return nil, rpcValidationError(err)
	}

	var nickname string
	if req.GetNickname() != nil {
		nickname = req.GetNickname().GetValue()
	}

	return &pb.UpdateUserResponse{
		UserUuid: "123e4567-e89b-12d3-a456-426614174000",
		Nickname: nickname,
	}, nil
}

func (s *server) SearchUsersByNickname(_ context.Context, req *pb.SearchUsersByNicknameRequest) (*pb.SearchUsersByNicknameResponse, error) {
	log.Printf("SearchUsersByNickname request sent with query %q", req.GetQuery())
	if err := s.validator.Validate(req); err != nil {
		return nil, rpcValidationError(err)
	}

	return &pb.SearchUsersByNicknameResponse{
		Users: []*pb.User{
			{
				UserUuid:    "11111111-1111-1111-1111-111111111111",
				Nickname:    "vasyan123",
				Age:         27,
				BirthDate:   timestamppb.New(time.Date(1997, 3, 15, 0, 0, 0, 0, time.UTC)),
				AvatarUrl:   "https://example.com/avatars/vasyan123.png",
				CreatedAt:   timestamppb.New(time.Date(2023, 5, 10, 14, 30, 0, 0, time.UTC)),
				SoftDeleted: false,
			},
			{
				UserUuid:    "22222222-2222-2222-2222-222222222222",
				Nickname:    "super_oleg",
				Age:         34,
				BirthDate:   timestamppb.New(time.Date(1990, 11, 2, 0, 0, 0, 0, time.UTC)),
				AvatarUrl:   "https://example.com/avatars/superoleg.jpg",
				CreatedAt:   timestamppb.New(time.Date(2022, 8, 20, 9, 15, 0, 0, time.UTC)),
				SoftDeleted: true,
			},
		},
	}, nil
}
