package server

import (
	"api-channel/proto"
	"context"
)

func (s *Server) AddMember(ctx context.Context, req *proto.AddMemberRequest) (*proto.AddMemberResponse, error) {

	for _, u := range req.MembersId {
		err := s.db.CreateChatMember(req.ChatId, u, "member")
		if err != nil {
			return nil, err
		}
	}
	return &proto.AddMemberResponse{}, nil
}
