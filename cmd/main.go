package main

import "github.com/sandronister/websocket-go/internal/infra/web"

func main() {

	s := web.NewServer("8080")
	s.Websocket()
	err := s.Run()

	if err != nil {
		panic(err)
	}
}
