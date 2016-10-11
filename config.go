package gouken

/*
config.go is decoupled from app.go
*/

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/spf13/viper"
)

// defineConfig define config for options
func defineConfig() {
	viper.SetDefault("name", "")
	viper.SetDefault("host", "localhost")
	viper.SetDefault("port", 20000)
	viper.SetDefault("client_address", "")
	viper.SetDefault("log_level", "INFO")
	viper.SetDefault("log_filename", false)
	viper.SetDefault("log_request", false)
	viper.SetDefault("log_response", false)
	viper.SetDefault("debug", true)
}

func confName() string {
	return viper.GetString("name")
}

func confHost() string {
	return viper.GetString("host")
}

func confPort() int {
	return viper.GetInt("port")
}

func confClientAddress() string {
	return viper.GetString("client_address")
}

func confLogLevel() string {
	return viper.GetString("log_level")
}

func confLogFilename() bool {
	return viper.GetBool("log_filename")
}

func confLogRequest() bool {
	return viper.GetBool("log_request")
}

func confLogResponse() bool {
	return viper.GetBool("log_response")
}

func confDebug() bool {
	return viper.GetBool("debug")
}

type ConfOption func()

func MakeConfig(filename string, opts ...ConfOption) {
	sp := strings.Split(filename, ".")
	if len(sp) != 2 {
		log.Fatalf("Could not parse config filename correctly: %v", filename)
	}

	viper.SetConfigName(sp[0])
	viper.SetConfigType(sp[1])

	// add cwd for config path
	viper.AddConfigPath(".")

	for _, o := range opts {
		o()
	}

	// read config at last
	ReadConfig()
}

func ConfPathEnv(n string) ConfOption {
	return func() {
		p := os.Getenv(n)
		log.Printf("%v: %v\n", n, p)
		if p != "" {
			viper.AddConfigPath(p)
		}
	}
}

func ConfEnvPrefix(n string) ConfOption {
	return func() {
		viper.SetEnvPrefix(n)
	}
}

func ConfBindEnv(n string) ConfOption {
	return func() {
		viper.BindEnv(n)
	}
}

func ConfNew(k string, v interface{}) ConfOption {
	return func() {
		viper.SetDefault(k, v)
	}
}

func ReadConfig() {
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func PrintConfig() {
	log.Printf("Config:\n%v", GetConfigString("  "))
}

// GetConfigString format all configs to string
func GetConfigString(prefix string) string {
	s := ""
	items := viper.AllSettings()
	keys := make([]string, len(items))

	// Sort keys
	i := 0
	for k := range items {
		keys[i] = k
		i++
	}
	sort.Strings(keys)

	for _, k := range keys {
		s += fmt.Sprintf("%v%v: %v\n", prefix, k, items[k])
	}
	return s
}
