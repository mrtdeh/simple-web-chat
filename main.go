package main

import (
	server "api-channel/pkg/server"
	"flag"
	"log"
	"time"
)

func main() {

	port := flag.String("p", "8082", "The port of server")
	flag.Parse()

	hostPort := "localhost:" + *port

	server := &server.Server{
		Addr: hostPort,
		Options: &server.ServerOptions{
			MaxTimeout: time.Second * 10,
		},
	}

	if err := server.Serve(); err != nil {
		log.Fatal(err)
	}
}
