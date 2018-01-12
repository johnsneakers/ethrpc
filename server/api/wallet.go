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

func (s *Server) Balance(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	address := r.Form.Get("address")
	ret,err := s.Client.EthGetBalance(address,"latest")
	if err != nil {
		panic(err)
	}

	fmt.Println("余额:", ret)
}

func (s *Server) CreateAccount(w http.ResponseWriter, r *http.Request) {
	fmt.Println("api..")
}
