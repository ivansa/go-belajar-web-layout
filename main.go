package main

import (
	"go-belajar-web-layout/handler"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/layout-file", http.HandlerFunc(handler.TemplateLayout))
	mux.Handle("/layout-define", http.HandlerFunc(handler.TemplateLayoutDefine))

	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
