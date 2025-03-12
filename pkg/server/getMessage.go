package server

import (
	"api-channel/proto"
	"time"
)

func (s *Server) GetMessages(req *proto.GetMessagesRequest, stream proto.ChatService_GetMessagesServer) error {
	time.Sleep(time.Second * 1)
	messages, err := s.db.GetMessages(req.ChatId, req.ReadedMsgId, req.Direction, req.Count)
	if err != nil {
		return err
	}

	lastID, err := s.db.GetLastMessageID(req.ChatId)
	if err != nil {
		return err
	}

	var res = &proto.MessagesResponse{}

	if messages[len(messages)-1].ID == lastID {
		res.Follow = true
	}

	var data []*proto.MessageData
	for _, m := range messages {

		// Fetch message attachment's placeholder
		var attachs []*proto.Attachment
		for _, att := range m.Attachments {
			for _, t := range att.Thumbnails {
				if t.Type == "placeholder" {
					attachs = append(attachs, &proto.Attachment{
						Placeholder: t.Base64,     // Base64 of the attachment placeholder
						Type:        att.FileType, // NOTE: most be image|video|audio|other...
						Url:         att.FilePath, // TODO: replace with http URL
					})
					break
				}
			}
		}

		// Fetch replied messages with thumbnails
		var repliedMessages []*proto.RepliedMessage
		for _, r := range m.Replies {
			var thumbnails []string
			for _, t := range r.Thumbnails {
				thumbnails = append(thumbnails, t.Thumbnail.Base64)
			}
			repliedMessages = append(repliedMessages, &proto.RepliedMessage{
				MessageId:  uint32(r.ReplyMessageId),
				Content:    r.ReplyMessage.Content,
				Thumbnails: thumbnails,
			})
		}

		// Collect message data
		data = append(data, &proto.MessageData{
			SenderId:        uint32(m.Sender.ID),
			MessageId:       uint32(m.ID),
			Content:         m.Content,
			SendAt:          m.CreatedAt.Unix(),
			Attachements:    attachs,
			RepliedMessages: repliedMessages,
		})

	}

	res.Data = data
	err = stream.Send(res)
	if err != nil {
		return err
	}

	return nil
}
