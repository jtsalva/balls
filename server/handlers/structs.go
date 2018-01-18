package handlers

type LoginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Success bool   `json:"success"`
}

type UpdateRequest struct {
	PlayerId   int    `json:"playerId"`
	PlayerName string `json:"playerName"`
	LobbyId    int    `json:"lobbyId"`
	MousePos   [2]int `json:"mousePos"`
}

type UpdateResponse struct {
	Updates []Update `json:"updates"`
}

type Update struct {
	Action   string                 `json:"action"`
	PlayerId int                    `json:"playerId"`
	Data     map[string]interface{} `json:"data"`
}

// type ConnectedPlayerInfo struct {
// 	PlayerId int `json:"playerId"`
// 	LobbyId  int `json:"lobbyId"`
// }
