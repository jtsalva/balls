package handlers

import (
	"encoding/json"
	"github.com/jtsalva/balls/game"
	// "github.com/jtsalva/balls/lobby"
	// "github.com/jtsalva/balls/player"
	"log"
	"net"
	"time"
)

func GameInstanceHandler() {

	// l := lobby.PlayerLobbies[lobby.LobbyIdToIndex(lobbyId)]
	// log.Println(l.OpeningInfo())

	// var updatedPlayers map[int]player.ConnectedPlayer

	// 60 whole calculations per second
	ticker := time.NewTicker(time.Second / 60)
	quit := make(chan struct{})

	// Keep running updates
	for {
		select {
		case <-ticker.C:
			// updatedPlayers = game.UpdatePlayers(l.ConnectedPlayers)
			// if p, ok := updatedPlayers[1234]; ok {
			// 	log.Println("Ballo", p.BallPos)
			// }
			game.Instance.UpdatePlayers()
			if game.Instance.Quit {
				close(quit)
			}
		case <-quit:
			ticker.Stop()
			break
		}
	}

	// log.Println(l.ClosingInfo())
}

func GameConnectionHandler(conn *net.UDPConn, portOut string) {
	log.Println("Game Connection Handler Started")

	buf := make([]byte, 1024)
	var syncRequest UpdateRequest

	for {
		if n, rAddr, err := conn.ReadFromUDP(buf); err != nil {
			log.Println(err)
		} else {
			if err := json.Unmarshal(buf[:n], &syncRequest); err != nil {
				log.Println(err)
			}

			log.Println(syncRequest.MousePos)

			// go syncWithLobby(syncRequest)
			game.Instance.ConnectPlayer(syncRequest.PlayerId, syncRequest.PlayerName, syncRequest.MousePos)
			syncWithGameInstance(syncRequest)
			go responseUpdate(conn, rAddr, syncRequest, portOut)

			// 	ch <- ConnectedPlayerInfo{
			// 		PlayerId: syncRequest.PlayerId,
			// 		LobbyId:  syncRequest.LobbyId}
		}
	}
	log.Println("Game Connection Handler Stopped")
}

func syncWithGameInstance(req UpdateRequest) {
	// log.Println("syncing")
	game.Instance.SyncPlayerMouse(req.PlayerId, req.MousePos)
}

// func syncWithLobby(req UpdateRequest) {
// 	lobby.PlayerLobbies[lobby.LobbyIdToIndex(req.LobbyId)].
// 		UpdatePlayerMousePos(req.PlayerId, req.MousePos)
// }

func responseUpdate(conn *net.UDPConn, addr *net.UDPAddr, req UpdateRequest, portOut string) {
	var updateResponse UpdateResponse
	data := make(map[string]interface{})

	// log.Println("legthy", len(lobbyPlayers))

	// log.Println("st")
	for _, player := range game.Instance.Players {
		// log.Println("in")
		data["pos"] = player.BallPos
		// data["lobbyId"] = req.LobbyId

		updateResponse.Updates = append(updateResponse.Updates, Update{
			Action:   "UPDATE_PLAYER",
			PlayerId: req.PlayerId,
			Data:     data})
	}

	if payload, err := json.Marshal(updateResponse); err != nil {
		log.Println(err)
	} else {
		addr.Port = 5671
		conn.WriteToUDP(payload, addr)
	}
}
