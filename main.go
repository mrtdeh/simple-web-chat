package main

import (
	server "api-channel/pkg/server"
	"flag"
	"log"
	"time"
)

func main() {

	port := flag.Int("p", 8082, "The port of server")
	flag.Parse()

	s := server.NewServer(server.Config{
		Port: *port,
		Options: server.ServerOptions{
			MaxTimeout: time.Second * 10,
		},
	})

	if err := s.Serve(); err != nil {
		log.Fatal(err)
	}
}
