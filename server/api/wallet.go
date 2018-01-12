package api

import (
	"fmt"
	"net/http"
	"ethrpc/server/utils"
	"math/big"
	"ethrpc"
	"strconv"
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

	data := map[string]interface{}{
		"balance":ret.String(),
	}

	utils.NewResp(w).RespSucc(data)
	return
}

func (s *Server) CreateAccount(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	pwd := r.Form.Get("pwd")
	ret,err := s.Client.Personal_newAccount(pwd)
	if err != nil {
		panic(err)
	}

	data := map[string]interface{}{
		"address":ret,
	}

	utils.NewResp(w).RespSucc(data)
	return
}


func (s *Server) Transaction(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	from := r.Form.Get("from")
	to := r.Form.Get("to")
	v := r.Form.Get("value")
	value_float,err := strconv.ParseFloat(v,10)
	if err != nil {
		panic(err)
	}

	//1000000000000000000 = 1 ETH
	vx := int64(value_float*1000000000000000000)
	t := ethrpc.T{
		From:     from,
		To:       to,
		Gas:      24900,
		GasPrice:  big.NewInt(5000000000),
		Value:    big.NewInt(vx),
		Data:     "0x61626364656667",
	}


	ps := ethrpc.PersonalCall{}

	p := t.ConverParam()
	ps.P = p
	ps.Pwd = "johnsneakers"
	ret,err := s.Client.Call("personal_sendTransaction", p)
	if err != nil {
		panic(err)
	}

	data := map[string]interface{}{
		"ret":ret,
	}

	utils.NewResp(w).RespSucc(data)


}