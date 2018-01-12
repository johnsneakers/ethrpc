package utils

import (
	"encoding/json"
	"net/http"
	"time"
)

var jsonContentType = []string{"application/json; charset=utf-8"}

type Resp struct {
	w http.ResponseWriter
}

func NewResp(w http.ResponseWriter) *Resp {
	return &Resp{w:w}
}

func (this *Resp)RespSucc(data map[string]interface{})  {
	this.w.Header()["Content-Type"] = jsonContentType
	ret := map[string]interface{}{
		"code": 1,
		"data": data,
	}

	jsonBytes, err := json.Marshal(ret)
	if err != nil {
		panic(err)
	}

	this.w.Write(jsonBytes)
}

func (this *Resp) RespError(data map[string]interface{}) {
	this.w.Header()["Content-Type"] = jsonContentType
	ret := map[string]interface{}{
		"code": -200,
		"data": data,
	}

	jsonBytes, err := json.Marshal(ret)
	if err != nil {
		panic(err)
	}

	this.w.Write(jsonBytes)
}

func Json(obj interface{}, w http.ResponseWriter) string {
	w.Header()["Content-Type"] = jsonContentType

	ret := map[string]interface{}{
		"_t": time.Now().Unix(),
		"_d": obj,
	}

	jsonBytes, err := json.Marshal(ret)
	if err != nil {
		panic(err)
	}
	w.Write(jsonBytes)
	return string(jsonBytes)
}

func WriteJson(obj interface{}, w http.ResponseWriter) {
	w.Header()["Content-Type"] = jsonContentType
	jsonBytes, _ := json.Marshal(obj)
	w.Write(jsonBytes)
}

func WriteJsonIndent(obj interface{}, w http.ResponseWriter) {
	w.Header()["Content-Type"] = jsonContentType
	jsonBytes, _ := json.MarshalIndent(obj, "  ", "    ")
	w.Write(jsonBytes)
}
