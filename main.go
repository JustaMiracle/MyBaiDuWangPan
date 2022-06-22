package main

import (
	"fileStoreServer/handler"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/file/upload", handler.UploadHandler)
	http.HandleFunc("/file/upload/suc", handler.UploadScHandler)
	http.HandleFunc("/file/meta", handler.GetFileHander)
	http.HandleFunc("/user/signup", handler.SiginUpHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Failed to start server ,err=%s", err.Error())
	}
}
