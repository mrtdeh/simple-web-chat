package server

import (
	database "api-channel/pkg/db"
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

	db := database.GetInstance()

	// Insert message into messages table
	msg := models.Message{
		SenderID: t.UserID,
		ChatID:   uint(req.ChatId),
		Content:  req.Content,
	}
	err = db.Create(&msg).Error
	if err != nil {
		return nil, err
	}

	err = db.Model(&models.Chat{}).Where("id = ?", req.ChatId).Update("last_message_id", msg.ID).Error
	if err != nil {
		return nil, err
	}

	for _, r := range req.RepliedMessages {
		// Create replied message
		replied := &models.Reply{
			MessageID:      msg.ID,
			ReplyMessageId: uint(r.MessageId),
		}
		db.Create(replied)
		// Create thumbnails related to replied message
		for _, sel := range r.SelectedAttachments {
			// Find mini thumbnail of attachment
			var tumbId uint
			db.Where(
				&models.Thumbnail{AttachmentID: uint(sel), Type: "mini"},
			).Select("id").First(&tumbId)
			// Create replied message thumbnail
			db.Create(&models.ReplyThumbnails{
				ReplyID:     replied.ID,
				ThumbnailId: tumbId,
			})
		}
	}

	return &proto.MessageResponse{MessageId: uint32(msg.ID)}, nil
}
