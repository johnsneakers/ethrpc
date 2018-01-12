package api

import (
	"net/http"

	"ethrpc/server/proto"
	"ethrpc/server/errorx"
	"ethrpc/server/utils"
)

type Server struct {
	Conf     *proto.Conf

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
					panic(err)
					utils.Json(map[string]int{"ret": -500}, w)
					return
				}
			}
		}()
		r.ParseForm()
		n.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
