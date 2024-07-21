package user

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (handler *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", handler.handleLogin).Methods("POST")
	router.HandleFunc("/register", handler.handleRegister).Methods("POST")
}

func (handler *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

}

func (handler *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {

}
