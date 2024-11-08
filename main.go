package main

import (
	database "api-channel/pkg/db"
	server "api-channel/pkg/server"
	"flag"
	"log"
	"time"
)

func main() {

	port := flag.Int("p", 9090, "The port of server")
	flag.Parse()

	db, err := database.New()
	if err != nil {
		log.Fatal(err)
	}

	s := server.NewServer(server.Config{
		Database: db,
		Port:     *port,
		Options: server.ServerOptions{
			MaxTimeout: time.Second * 10,
		},
	})

	if err := s.Serve(); err != nil {
		log.Fatal(err)
	}
}
