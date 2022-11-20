package main

import (
	"os"

	"github.com/seki-shinnosuke/study-golang/config"
	"github.com/seki-shinnosuke/study-golang/server"
	"github.com/seki-shinnosuke/study-golang/util/logger"
)

func main() {
	c := config.NewConfig("app.env")
	os.Setenv("TZ", c.TimeZone)

	db, err := config.NewDB(c.TimeZone, &c.RDB)
	if err != nil {
		logger.Fatal("Failed to Connect Database. err: %v", err)
	}

	s := server.InitializeService(&c.APIServer, db)
	if err := s.Run(); err != nil {
		logger.Fatal("Failed to start API server. err: %v", err)
	}
}
