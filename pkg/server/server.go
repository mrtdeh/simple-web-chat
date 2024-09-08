package server

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type ServerOptions struct {
	MaxTimeout           time.Duration
	MaxConnectTry        int
	MaxConnectTryTimeout time.Duration
}
type Server struct {
	Id      string
	Addr    string
	status  string
	Options *ServerOptions
}

func (a *Server) Serve() error {
	a.setDefaultOptions()

	if a.Id == "" {
		id := sha256.Sum256([]byte(time.Now().String()))
		serverId := hex.EncodeToString(id[:])
		a.Id = serverId
	}
	a.status = "unknown"
	chatServer := NewChat(a.Id)

	return chatServer.Serve(a.Addr)
}

func (a *Server) setDefaultOptions() {
	if a.Options == nil {
		a.Options = &ServerOptions{}
	}
	if a.Options.MaxConnectTry == 0 {
		a.Options.MaxConnectTry = 10
	}
	if a.Options.MaxTimeout == 0 {
		a.Options.MaxTimeout = time.Second * 5
	}
	if a.Options.MaxConnectTryTimeout == 0 {
		a.Options.MaxConnectTryTimeout = time.Second * 1
	}
}

func (a *Server) Status() string {
	return a.status
}
