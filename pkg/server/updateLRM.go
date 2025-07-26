package server

import (
	"api-channel/proto"
	"context"
)

func (s *Server) UpdateLastReadedMessageID(ctx context.Context, req *proto.LrmRequest) (*proto.Close, error) {
	t, err := checkToken(req.Token)
	if err != nil {
		return nil, err
	}

	err = s.db.SetChatLastReadedMessageID(req.ChatId, t.UserID, req.ReadedMsgId)
	if err != nil {
		return nil, err
	}

	// s.noticeUser(t.Username)
	return &proto.Close{}, nil
}
