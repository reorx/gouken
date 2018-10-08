package gouken

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

type Config struct {
	Name        string
	Host        string
	Port        int
	Logger      *logrus.Logger
	LogFilename bool
	LogRequest  bool
	LogResponse bool
	Debug       bool
}

func (c Config) addr() string {
	return fmt.Sprintf("%v:%v", c.Host, c.Port)
}

func (c Config) Check() error {
	if c.Logger == nil {
		return fmt.Errorf("config's logger shoule not be nil")
	}
	return nil
}
