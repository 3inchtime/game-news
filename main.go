package main

import (
	"game-news/routers"
	"net/http"
)

func main() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr:    "0.0.0.0:8000",
		Handler: router,
	}

	s.ListenAndServe()
}
