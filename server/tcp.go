package server

import (
	"github.com/gorilla/mux"
	"github.com/jtsalva/balls/server/handlers"
	"log"
	"net/http"
)

func StartTCPServer(port string) {
	log.Println("Starting TCP Server")

	r := mux.NewRouter()

	r.HandleFunc("/login", handlers.LoginHandler)

	http.ListenAndServe(port, r)
}
