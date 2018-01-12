package api

import (
	"net/http"

	"ethrpc/server/proto"
	"ethrpc/server/errorx"
	"ethrpc/server/utils"
	"ethrpc"
)

type Server struct {
	Conf     *proto.Conf
	Client *ethrpc.EthRPC
}

func (s *Server) RecoverHandler(n http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				switch err.(type) {
				case *errorx.Error:
					e := err.(*errorx.Error)
					j, _ := e.MarshalJSON()
					w.Header()["Content-Type"] = []string{"application/json; charset=utf-8"}
					w.Write(j)
					return
				default:
					utils.NewResp(w).RespError(map[string]interface{}{"code": -500, "msg":err})
					return
				}
			}
		}()
		r.ParseForm()
		n.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
