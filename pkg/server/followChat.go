package server

import (
	"api-channel/proto"
)

func (s *Server) FollowChat(req *proto.FollowChatRequest, stream proto.ChatService_FollowChatServer) error {

	lastID, err := s.db.GetLastMessageID(req.ChatId)
	if err != nil {
		return err
	}

	if lastID == req.LastMsgId {

	}

	return nil
}
