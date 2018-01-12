package main

import (
	"fmt"
	"ethrpc/server/proto"
	"log"
	"net/http"
	"github.com/justinas/alice"
	"ethrpc/server/api"
	"ethrpc"
)

func main() {
	servicesConf, err := proto.LoadServiceConf()
	if err != nil {
		log.Panicln(err)
	}

	s := new(api.Server)
	s.Client = ethrpc.NewEthRPC("http://127.0.0.1:8545")
	common := alice.New(s.RecoverHandler)
	http.Handle("/account/create", common.ThenFunc(s.CreateWallet))
	http.Handle("/debug", common.ThenFunc(s.VersionCheck))
	RunHttpServer(servicesConf)
}


func RunHttpServer(conf *proto.Conf) {
	addr := fmt.Sprintf("%s:%d", conf.Server.Host, conf.Server.Port)
	log.Printf("Starting Ethereum json rpc Version:%s Port:%d", conf.Server.Version, conf.Server.Port)
	if e := http.ListenAndServe(addr, nil); e != nil {
		log.Panicln(e)
	}
}
