package main

import "net/http"

type APIServer struct {
	listenAddr string
}

func NewAPIServer(listenAddr string) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
	}
}

func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) {

}

func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) {

}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) {

}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) {

}

func (s *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) {

}
