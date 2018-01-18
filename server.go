package main

import (
	// "github.com/jtsalva/balls/game"
	"github.com/jtsalva/balls/server"
	"github.com/jtsalva/balls/server/handlers"
)

func StartServers() {
	// Global main lobby

	go server.StartTCPServer(":80")
	go handlers.GameInstanceHandler()
	server.StartGameServer(server.UDPConfig{
		PortIn:  ":5672",
		PortOut: ":5671"})
}
