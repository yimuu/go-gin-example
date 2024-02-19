package setting

import (
	"github.com/spf13/viper"
	"log"
	"time"
)

var (
	ViperConfig *viper.Viper

	RunMode string

	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	PageSize  int
	JwtSecret string
)

func init() {
	// Load config file
	ViperConfig = viper.New()
	ViperConfig.SetConfigFile("conf/app.toml")
	err := ViperConfig.ReadInConfig()
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.toml': %ViperConfig", err)
	}

	// Set config
	RunMode = ViperConfig.GetString("runMode")
	HttpPort = ViperConfig.GetInt("server.httpPort")
	ReadTimeout = time.Duration(ViperConfig.GetInt("server.readTimeout")) * time.Second
	WriteTimeout = time.Duration(ViperConfig.GetInt("server.writeTimeout")) * time.Second
	PageSize = ViperConfig.GetInt("app.pageSize")
	JwtSecret = ViperConfig.GetString("app.jwtSecret")
}
