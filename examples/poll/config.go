package poll

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

// InitConfig ..
func InitConfig() {
	// define
	viper.SetDefault("log_level", "INFO")
	viper.SetDefault("host", "localhost")
	viper.SetDefault("port", 9000)
	viper.SetDefault("debug", false)

	// env
	viper.SetEnvPrefix("poll")
	viper.BindEnv("log_level")

	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")

	projectpath := os.Getenv("PROJECTPATH")
	log.Println("PROJECTPATH:", projectpath)
	if projectpath != "" {
		viper.AddConfigPath(projectpath)
	}

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
