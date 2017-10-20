package api

import (
	"fmt"
	"ferp-api/pkg/model"
	"github.com/ant0ine/go-json-rest/rest"
)

func register(w rest.ResponseWriter, r *rest.Request) {
	body := map[string]string{}
	err := r.DecodeJsonPayload(&body)
	if err != nil {
		w.WriteHeader(400)
		w.WriteJson(err)
		return
	}
	resp, status := model.RegisterUser(body)
	w.WriteHeader(status)
	w.WriteJson(resp)
}

func upload(w rest.ResponseWriter, r *rest.Request) {
	fmt.Println("test")
	w.WriteHeader(200)
	w.WriteJson("ok")
}