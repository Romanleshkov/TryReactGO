package main

import (
	"backend/api"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	var handlers api.Handler
	srv := new(api.Server)
	if err := srv.Run(port, handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occurred while running http server: %s", err.Error())
	}
}
