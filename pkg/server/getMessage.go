package server

import (
	"api-channel/proto"
	"fmt"
	"time"
)

func (s *Server) GetMessages(req *proto.GetMessagesRequest, stream proto.ChatService_GetMessagesServer) error {
	time.Sleep(time.Second * 1)

	t, err := checkToken(req.Token)
	if err != nil {
		return err
	}

	sess := s.Sessions.Get(t.Username)
	if sess == nil {
		return fmt.Errorf("session for '%s' is expired or not found", t.Username)
	}
	// sess.sl.Lock()
	// defer sess.sl.Unlock()

	sess.activeChat(req.ChatId)
	sess.unfollow()

	if req.PageMax == 0 {
		return fmt.Errorf("page_max must be exist")
	}

	lastID, err := s.db.GetLastMessageID(req.ChatId)
	if err != nil {
		return err
	}

	if req.Direction == proto.GetMessagesRequest_BothPage {
		if req.FromMsgId == 0 {
			lastReadedId, err := s.db.GetChatLastReadedMessageID(req.ChatId, t.UserID)
			if err != nil {
				return err
			}
			req.FromMsgId = lastReadedId
		}
	}

	if req.Direction == proto.GetMessagesRequest_None {

		if req.LastMsgId > 0 { // If lastMsgId is exist then check with real last message id from db
			if req.LastMsgId == lastID {
				sess.follow()
			}
			// TODO : may need to send some think
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
			sess.follow()
		}
		if req.Direction == proto.GetMessagesRequest_PrevPage {
			if req.LastMsgId == lastID {
				sess.follow()
			}
		}

	} else {
		// else if messages available then in prev calculate page size to determine is last page or not
		// and in other direction check last message id.

		if req.Direction == proto.GetMessagesRequest_PrevPage {
			msgsLen := len(messages)
			a := req.PageSize + int32(msgsLen)
			if a < req.PageMax {
				sess.follow()
			}
		} else {
			if messages[len(messages)-1].ID == lastID {
				sess.follow()
			}
		}
	}

	protoMessages := convertMessagesToProto(messages)
	err = stream.Send(protoMessages)
	if err != nil {
		return err
	}

	return nil
}
