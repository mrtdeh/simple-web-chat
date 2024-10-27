package server

import (
	database "api-channel/pkg/db"
	"api-channel/pkg/models"
	"api-channel/proto"
	"context"
)

func (s *Server) AddUser(ctx context.Context, req *proto.AddUserRequest) (*proto.AddUserResponse, error) {

	db := database.GetInstance()
	// TODO: check username and password in db
	var user = models.User{
		Username:    req.Username,
		Password:    req.Password,
		PhoneNumber: req.PhoneNumber,
		Email:       req.Email,
		DisplayName: req.DisplayName,
		// ProfilePicture: req.,
	}
	res := db.Create(&user)
	if res.Error != nil {
		return nil, res.Error
	}

	return &proto.AddUserResponse{UserId: uint32(user.ID)}, nil

}
