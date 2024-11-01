package server

import (
	"api-channel/pkg/models"
	"api-channel/proto"
	"context"
)

func (s *Server) AddUser(ctx context.Context, req *proto.AddUserRequest) (*proto.AddUserResponse, error) {

	userId, err := s.db.CreateUser(models.User{
		Username:       req.Username,
		Password:       req.Password,
		DisplayName:    req.DisplayName,
		PhoneNumber:    req.PhoneNumber,
		Email:          req.Email,
		ProfilePicture: "",
		Bio:            "",
	})
	if err != nil {
		return nil, err
	}

	return &proto.AddUserResponse{UserId: uint32(userId)}, nil
}
