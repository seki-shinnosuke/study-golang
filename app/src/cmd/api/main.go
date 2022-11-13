package main

import (
	"os"

	"github.com/seki-shinnosuke/study-golang/config"
	"github.com/seki-shinnosuke/study-golang/server"
	"github.com/seki-shinnosuke/study-golang/util/logger"
)

func main() {
	os.Setenv("TZ", "Asia/Tokyo")
	c := config.NewConfig("app.env")
	s := server.InitializeService(&c.APIServer)
	if err := s.Run(); err != nil {
		logger.Fatal("Failed to start API server. err: %v", err)
	}
}
