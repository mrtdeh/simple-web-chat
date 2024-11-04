package server

import (
	"api-channel/proto"
	"context"
)

func (s *Server) GetMessages(ctx context.Context, req *proto.GetMessagesRequest) (*proto.MessagesResponse, error) {

	// return &proto.Messages{Data: data}, nil
	messages, err := s.db.GetMessages(req.ChatId)
	if err != nil {
		return nil, err
	}

	var data []*proto.MessagesResponse_MessageData
	for _, m := range messages {

		// Fetch message attachment's placeholder
		// var attachsPlaceholder []string
		var attachs []*proto.MessagesResponse_Attachment
		for _, att := range m.Attachments {
			for _, t := range att.Thumbnails {
				if t.Type == "placeholder" {
					attachs = append(attachs, &proto.MessagesResponse_Attachment{
						Placeholder: t.Base64,     // Base64 of the attachment placeholder
						Type:        att.FileType, // NOTE: most be image|video|audio|other...
						Url:         att.FilePath, // TODO: replace with http URL
					})
					break
				}
			}
		}

		// Fetch replied messages with thumbnails
		var repliedMessages []*proto.MessagesResponse_RepliedMessage
		for _, r := range m.Replies {
			var thumbnails []string
			for _, t := range r.Thumbnails {
				thumbnails = append(thumbnails, t.Thumbnail.Base64)
			}
			repliedMessages = append(repliedMessages, &proto.MessagesResponse_RepliedMessage{
				MessageId:  uint32(r.ReplyMessageId),
				Content:    r.ReplyMessage.Content,
				Thumbnails: thumbnails,
			})
		}

		// Collect message data
		data = append(data, &proto.MessagesResponse_MessageData{
			SenderId:        uint32(m.Sender.ID),
			MessageId:       uint32(m.ID),
			Content:         m.Content,
			SendAt:          m.CreatedAt.Unix(),
			Attachements:    attachs,
			RepliedMessages: repliedMessages,
		})

	}

	return &proto.MessagesResponse{Data: data}, nil
}
