package main

import (
	"os"

	"github.com/seki-shinnosuke/study-golang/config"
	"github.com/seki-shinnosuke/study-golang/server"
	"github.com/seki-shinnosuke/study-golang/util/logger"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func main() {
	c := config.NewConfig("app.env")
	os.Setenv("TZ", c.TimeZone)

	db, err := config.NewDB(c.TimeZone, &c.RDB)
	if err != nil {
		logger.Fatal("Failed to Connect Database. err: %v", err)
	}
	boil.SetDB(db)

	s := server.InitializeService(&c.APIServer, db)
	if err := s.Run(); err != nil {
		logger.Fatal("Failed to start API server. err: %v", err)
	}
}
