package server

import (
	// "github.com/jtsalva/balls/game"
	"github.com/jtsalva/balls/server/handlers"
	"log"
	"net"
	"time"
)

type UDPConfig struct {
	PortIn  string `json:"portIn"`
	PortOut string `json:"portOut"`
}

type Connection struct {
	PlayerId       int       `json:"playerId"`
	LobbyId        int       `json:"lobbyId"`
	LastConnection time.Time `json:"lastConnection"`
}

func StartGameServer(config UDPConfig) {
	log.Println("Starting Game Server")
	protocol := "udp"

	lAddr, err := net.ResolveUDPAddr(protocol, config.PortIn)
	if err != nil {
		log.Println(err)
	}

	conn, err := net.ListenUDP(protocol, lAddr)
	if err != nil {
		log.Println(err)
	}

	// ch := make(chan handlers.ConnectedPlayerInfo)

	// Mapping playerId to Connection
	// connectedPlayers := make(map[int]Connection)

	// Start two connection handlers
	handlers.GameConnectionHandler(conn, config.PortOut)
	// go handlers.GameConnectionHandler(conn, config.PortOut)

	// Handle channel data
	// go func(ch chan handlers.ConnectedPlayerInfo) {
	// 	for data := range ch {
	// 		connectedPlayers[data.PlayerId] = Connection{
	// 			PlayerId:       data.PlayerId,
	// 			LobbyId:        data.LobbyId,
	// 			LastConnection: time.Now()}
	// 	}
	// }(ch)
}
