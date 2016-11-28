package main

import (
	"log"
	"time"

	"github.com/reorx/gouken/examples/poll"
)

func main() {
	app := poll.App()
	app.PrintConfig()
	app.OnStop(func() error {
		log.Println("call stop callback")
		return nil
	})
	go app.Run()

	time.Sleep(time.Duration(3) * time.Second)
	app.Stop()
}
