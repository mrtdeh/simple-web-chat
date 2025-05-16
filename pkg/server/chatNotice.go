package server

import (
	"api-channel/proto"
	"context"
)

func (s *Server) ChatNotice(ctx context.Context, req *proto.ChatNoticeRequest) (*proto.Close, error) {
	t, err := checkToken(req.Token)
	if err != nil {
		return nil, err
	}

	err = s.db.SetChatLastReadedMessageID(req.ChatId, t.UserID, req.ReadedMsgId)
	if err != nil {
		return nil, err
	}

	return &proto.Close{}, nil
}
