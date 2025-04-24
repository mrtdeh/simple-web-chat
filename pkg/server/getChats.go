package server

import (
	"api-channel/proto"
	"context"
	"log"
)

func (s *Server) GetChats(ctx context.Context, req *proto.GetChatsRequest) (*proto.Chats, error) {

	t, err := checkToken(req.Token)
	if err != nil {
		return nil, err
	}

	chats, err := s.getChats(t.UserID)
	if err != nil {
		return nil, err
	}

	return &proto.Chats{Data: chats}, nil
}

func (s *Server) getChats(userId uint32) ([]*proto.ChatData, error) {

	rows, err := s.db.GetChats(userId)
	if err != nil {
		return nil, err
	}

	var Data []*proto.ChatData
	for rows.Next() {
		var r *proto.ChatData
		if err := s.db.GORM().ScanRows(rows, &r); err != nil {
			log.Fatalf("failed to scan row: %v", err)
		}
		Data = append(Data, r)
	}

	return Data, nil
}
