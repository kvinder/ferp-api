package api

import (
	"ferp-api/pkg/model"
	"fmt"
	"mime/multipart"

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

func originFileName(mf *multipart.Form) string {
	return mf.File["file"][0].Filename
}

func openFile(data *multipart.Form) (multipart.File, error) {
	files := data.File["file"]
	file, err := files[0].Open()
	return file, err
}

func upload(w rest.ResponseWriter, r *rest.Request) {
	r.ParseMultipartForm(32 << 20)
	data := r.MultipartForm

	originalFilename := originFileName(data)
	fmt.Printf("% #v\n", originalFilename)
	file, err := openFile(data)
	if err != nil {
		fmt.Println(err)
	}
	defer func() { file.Close() }()
	fmt.Printf("% #v\n", file)

	w.WriteHeader(200)
	w.WriteJson(data)
}
