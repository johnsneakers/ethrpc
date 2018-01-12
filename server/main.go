package main

import (
	"fmt"
	"ethrpc/server/proto"
	"log"
	"net/http"
	"github.com/justinas/alice"
	"ethrpc/server/api"
)

func main() {
	servicesConf, err := proto.LoadServiceConf()
	if err != nil {
		log.Panicln(err)
	}

	s := new(api.Server)
	common := alice.New(s.RecoverHandler)
	http.Handle("/wallet/create", common.ThenFunc(s.CreateWallet))
	RunHttpServer(servicesConf)
}


func RunHttpServer(conf *proto.Conf) {
	addr := fmt.Sprintf("%s:%d", conf.Server.Host, conf.Server.Port)
	log.Printf("Starting Ethereum json rpc Version:%s Port:%d", conf.Server.Version, conf.Server.Port)
	if e := http.ListenAndServe(addr, nil); e != nil {
		log.Panicln(e)
	}
}
