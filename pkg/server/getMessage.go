package server

import (
	"api-channel/proto"
	"fmt"
	"time"
)

func (s *Server) GetMessages(req *proto.GetMessagesRequest, stream proto.ChatService_GetMessagesServer) error {
	if req.PageMax == 0 {
		return fmt.Errorf("page_max must be exist")
	}

	time.Sleep(time.Second * 1)
	var res = &proto.MessagesResponse{}

	lastID, err := s.db.GetLastMessageID(req.ChatId)
	if err != nil {
		return err
	}

	if req.Direction == proto.GetMessagesRequest_None {

		if req.LastMsgId > 0 { // If lastMsgId is exist then check with real last message id from db
			if req.LastMsgId == lastID {
				res.Follow = true
			}
			err = stream.Send(res)
			if err != nil {
				return err
			}
			return nil

		} else { // else, change direction of none to next for getting messages of begening from db
			req.Direction = proto.GetMessagesRequest_NextPage
			req.FromMsgId = 0
			if req.Count == 0 {
				req.Count = req.PageMax
			}
		}

	}

	messages, err := s.db.GetMessages(req.ChatId, req.FromMsgId, req.Direction, req.Count)
	if err != nil {
		return err
	}

	if len(messages) == 0 {
		// if not message available in next/last/both direction enable follow

		if req.Direction == proto.GetMessagesRequest_LastPage ||
			req.Direction == proto.GetMessagesRequest_NextPage ||
			req.Direction == proto.GetMessagesRequest_BothPage {
			res.Follow = true
		}
		if req.Direction == proto.GetMessagesRequest_PrevPage {
			if req.LastMsgId == lastID {
				res.Follow = true
			}
		}

	} else {
		// else if messages available then in prev calculate page size to determine is last page or not
		// and in other direction check last message id.

		if req.Direction == proto.GetMessagesRequest_PrevPage {
			msgsLen := len(messages)
			a := req.PageSize + int32(msgsLen)
			if a < req.PageMax {
				res.Follow = true
			}
		} else {
			if messages[len(messages)-1].ID == lastID {
				res.Follow = true
			}
		}
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
