package api

import (
	"fmt"
	"net/http"
)

func (s *Server) CreateWallet(w http.ResponseWriter, r *http.Request) {
	fmt.Println("api..")
}
