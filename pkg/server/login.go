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

	err := db.Where("username = ? AND password = ?", req.Username, req.Password).First(&result).Error
	if err != nil {
		return nil, fmt.Errorf("username or password incorrect")
	}

	// Create new token
	token := tm.Create(result.ID, result.Username)

	return &proto.LoginResponse{
		Token: token,
	}, nil
}
