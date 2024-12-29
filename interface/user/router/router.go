package router

import (
	"github.com/gorilla/mux"
	"github.com/yusuke-takatsu/go-training/interface/user/handler"
)

func NewRouter(handler *handler.Handler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/register", handler.Register).Methods("POST")

	return r
}
