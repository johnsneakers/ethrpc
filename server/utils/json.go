package utils

import (
	"encoding/json"
	"net/http"
	"time"
)

var jsonContentType = []string{"application/json; charset=utf-8"}

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
