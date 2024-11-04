package server

import (
	"api-channel/pkg/models"
	"api-channel/proto"
	"context"
	"fmt"
)

func (s *Server) UpdateMemberInfo(ctx context.Context, req *proto.MemberInfoRequest) (*proto.MemberInfoResponse, error) {
	t, err := checkToken(req.Token)
	if err != nil {
		return nil, err
	}

	chatMember := &models.ChatMember{
		UserID: t.UserID,
		ChatID: req.ChatId,
	}

	switch v := req.Info.(type) {
	case *proto.MemberInfoRequest_Mute:
		chatMember.Mute = v.Mute

	case *proto.MemberInfoRequest_LastReadedMessageId:
		chatMember.LastReadedMessageID = v.LastReadedMessageId

	default:
		return nil, fmt.Errorf("unsupported member info request")
	}

	err = s.db.GORM().Save(chatMember).Error
	if err != nil {
		return nil, err
	}

	return &proto.MemberInfoResponse{}, nil
}
