package api

import (
	"fmt"
	"net/http"
)


func (s *Server) VersionCheck(w http.ResponseWriter, r *http.Request) {
	str,err := s.Client.Web3ClientVersion()
	if err != nil {
		panic(err)
	}

	fmt.Println("version:", str)
}



func (s *Server) CreateWallet(w http.ResponseWriter, r *http.Request) {
	fmt.Println("api..")
}
