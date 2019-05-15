package main

import (
	"fmt"
	"game-news/routers"
	"net/http"
)

func main() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr:    "0.0.0.0:8000",
		Handler: router,
	}

	err := s.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
