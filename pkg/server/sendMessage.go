package server

import (
	"api-channel/pkg/models"
	"api-channel/proto"
	"context"
)

func (s *Server) SendMessage(ctx context.Context, req *proto.MessageRequest) (*proto.MessageResponse, error) {
	token := req.Token

	// Check token is exist
	t, err := checkToken(token)
	if err != nil {
		return nil, err
	}

	msgId, err := s.db.CreateMessage(req.ChatId, t.UserID, req.Content)
	if err != nil {
		return nil, err
	}

	for _, r := range req.RepliedMessages {
		// Create replied message
		replied := &models.Reply{
			MessageID:      msgId,
			ReplyMessageId: r.MessageId,
		}
		s.db.GORM().Create(replied)
		// Create thumbnails related to replied message
		for _, sel := range r.SelectedAttachments {
			// Find mini thumbnail of attachment
			var tumbId uint32
			s.db.GORM().
				Where(
					&models.Thumbnail{AttachmentID: sel, Type: "mini"},
				).Select("id").First(&tumbId)
			// Create replied message thumbnail
			s.db.GORM().
				Create(&models.ReplyThumbnails{
					ReplyID:     replied.ID,
					ThumbnailId: tumbId,
				})
		}
	}

	return &proto.MessageResponse{MessageId: msgId}, nil
}
