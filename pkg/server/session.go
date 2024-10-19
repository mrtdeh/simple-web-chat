package server

func (s *SessionManger) Add(key string, session *Session) {
	s.l.Lock()
	s.sessions[key] = session
	s.l.Unlock()
}

func (s *SessionManger) Delete(key string) {
	s.l.Lock()
	delete(s.sessions, key)
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

func (s *Session) FireReceive() {
	close(s.receive)
	s.receive = make(chan struct{})
}

func (s *Session) OnReceive() <-chan struct{} {
	return s.receive
}
