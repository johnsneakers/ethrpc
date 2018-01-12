package proto

import (
	"flag"
	"io/ioutil"
	"os"

	"github.com/naoina/toml"
)

type Conf struct {
	Server   *Server
}



type Server struct {
	Host          string
	Port          int
	LogDir        string
	Version       string
}

var (
	rootConfigPath = flag.String("c", "conf", "service root config path")
)

func LoadServiceConf() (conf *Conf, err error) {
	path := *rootConfigPath + "/" + "server.conf"
	file, err := os.Open(path)
	if err != nil {
		panic(err)
		return nil, err
	}
	defer file.Close()

	var buf []byte
	if buf, err = ioutil.ReadAll(file); err != nil {
		panic(err)
		return nil, err
	}

	conf = &Conf{}
	if err = toml.Unmarshal(buf, conf); err != nil {
		panic(err)
		return nil, err
	}
	return conf, err

}
