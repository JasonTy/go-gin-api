package main

import (
	"fmt"
	"net/http"

	"go/gin/api/pkg/setting"
	"go/gin/api/routers"

	"go/gin/api/models"
)

func main() {
	defer models.CloseDB()
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HttpPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeOut,
		WriteTimeout:   setting.WriteTimeOut,
		MaxHeaderBytes: 1 << 20,
	}


	s.ListenAndServe()
}