package main

import (
	"os"

	"github.com/seki-shinnosuke/study-golang/config"
	"github.com/seki-shinnosuke/study-golang/controller"
	"github.com/seki-shinnosuke/study-golang/logger"
)

func main() {
	os.Setenv("TZ", "Asia/Tokyo")
	c := config.NewConfig("app.env")
	r := controller.NewRouting(&c.APIServer)
	if err := r.Run(); err != nil {
		logger.Fatal("API起動エラー. err: %v", err)
	}
}
