package main

import (
	"game-news/routers"
	"net/http"
)

func main() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr:    "127.0.0.1:8000",
		Handler: router,
	}

	s.ListenAndServe()
}
