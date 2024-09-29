package server

import (
	database "api-channel/pkg/db"
	"api-channel/pkg/models"
	"api-channel/proto"
	"context"
	"fmt"
)

func (s *Server) Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {
	// user := req.Username
	// pass := req.Password

	db := database.GetInstance()
	// TODO: check username and password in db
	var result models.User

	db.Model(&models.User{
		Username: req.Username,
		Password: req.Password,
	}).Find(&result)
	if result.Username != req.Username {
		return nil, fmt.Errorf("username or password incorrect")
	}

	// Create new token
	token := tm.Create(result.ID, result.Username)

	return &proto.LoginResponse{
		Token: token,
	}, nil
}
