package server

import (
	"api-channel/pkg/models"
	"api-channel/proto"
	"fmt"
)

func (s *SessionManger) Add(key string, session *Session) {
	s.l.Lock()
	s.sessions[key] = session
	fmt.Println("session added : ", key)
	s.l.Unlock()
}

func (s *SessionManger) Delete(key string) {
	s.l.Lock()
	delete(s.sessions, key)
	fmt.Println("session deleted : ", key)
	s.l.Unlock()
}

func (s *SessionManger) Get(key string) *Session {
	s.l.RLock()
	defer s.l.RUnlock()
	if s, ok := s.sessions[key]; ok {
		return s
	}
	return nil
}
func (s *SessionManger) Close(key string) {
	s.l.Lock()
	defer s.l.Unlock()
	if s, ok := s.sessions[key]; ok {
		closeChannel(s.close)
	}
}

func (s *Session) follow() {
	s.folow = true
}

func (s *Session) unfollow() {
	s.folow = false
}

func (s *Session) send(m *proto.StreamResponse) error {
	s.sl.Lock()
	defer s.sl.Unlock()
	err := s.stream.Send(m)
	if err != nil {
		return err
	}
	return nil
}

// ===============================================
func convertMessagesToProto(messages []models.Message) *proto.Messages {
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
	return &proto.Messages{Data: data}
}

func (s *Session) sendMessages(messages []models.Message) error {
	s.sl.Lock()
	defer s.sl.Unlock()

	res := &proto.StreamResponse{}

	protoMessages := convertMessagesToProto(messages)

	res.Data = &proto.StreamResponse_Messages{
		Messages: protoMessages,
	}
	err := s.stream.Send(res)
	if err != nil {
		return err
	}

	return nil
}

func (s *Session) activeChat(chatId uint32) {
	s.activeChatId = chatId
}

func (s *Session) onReceive() <-chan *proto.StreamResponse {
	return s.receive
}
