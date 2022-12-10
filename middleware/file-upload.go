package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func UploadFile(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		file, handler, err := r.FormFile("input-image")

		if err != nil {
			fmt.Println(err)
			json.NewEncoder(w).Encode("Eror Retrieving the File.")
			return
		}

		fmt.Printf("Uploaded File: %v\n", handler.Filename)

		tempFile, err := ioutil.TempFile("uploads", "image-*"+handler.Filename)
		if err != nil {
			fmt.Println(err)
			fmt.Println("path upload eror.")
			json.NewEncoder(w).Encode(err)
			return
		}
		defer tempFile.Close()

		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
			return
		}
		tempFile.Write(fileBytes)

		data := tempFile.Name()
		filename := data[8:] // uploads/image-1237676812368wahyu.png (Mengurangi 8 huruf dari depan)

		ctx := context.WithValue(r.Context(), "dataFile", filename)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
