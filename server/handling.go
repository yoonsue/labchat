package server

import (
	"net/http"

	"github.com/gorilla/mux"

	kitlog "github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
)

func MakeHandler(hs Service, logger kitlog.Logger) http.Handler {
	r := mux.NewRouter()

	opts := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(logger),
		kithttp.ServerErrorEncoder(encodeError),
	}

	registerIncidentHandler := kithttp.NewServer(
		makeRegisterIncidentEndpoint(hs),
		decodeRegisterIncidentRequest,
		encodeResponse,
		opts...,
	)

	r.Handle("/keyboard", keyboardInit).Methods("GET")
	r.Handle("/message", message).Methods("POST")
	r.Handle("/friend", friend).Methods("POST")
	r.Handle("/friend/", removeFriend).Methods("DELETE")
	r.Handle("/chat_room/", leaveChatRoom).Methods("DELETE")

	return r
}
