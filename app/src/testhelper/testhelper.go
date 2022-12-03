package testhelper

import (
	"net/http/httptest"
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
	"github.com/seki-shinnosuke/study-golang/config"
	"github.com/seki-shinnosuke/study-golang/util/logger"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type TomlConfig struct {
	Mysql Mysql `toml:"mysql"`
}

type Mysql struct {
	Whitelist []string `toml:"whitelist"`
}

func NewTestConfig(t *testing.T) *config.Config {
	t.Helper()
	c := config.NewConfig("/app/src/test.app.env")
	return c
}

func SetupTestDB(t *testing.T) {
	t.Helper()

	c := NewTestConfig(t)
	db, err := config.NewDB(c.TimeZone, &c.RDB)
	if err != nil {
		logger.Fatal("Failed to Connect Database. err: %v", err)
	}
	boil.SetDB(db)
}

func CleanUpDB(t *testing.T) {
	t.Helper()

	var tomlConfig TomlConfig
	_, err := toml.DecodeFile("/database/sqlboiler/config.toml", &tomlConfig)
	if err != nil {
		logger.Fatal("Failed to Load Sqlboiler Config. err: %v", err)
	}

	db := boil.GetDB()
	for i := range tomlConfig.Mysql.Whitelist {
		_, err := db.Exec("TRUNCATE TABLE " + tomlConfig.Mysql.Whitelist[i])
		if err != nil {
			t.Error(err)
		}
	}
}

func NewHttp(t *testing.T) (*gin.Context, *gin.Engine, *httptest.ResponseRecorder) {
	t.Helper()

	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)

	return c, r, w
}
