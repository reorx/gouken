package main

import (
	"log"
	"time"

	"github.com/reorx/gouken/examples/poll"
)

func main() {
	app := poll.App()
	app.OnStop(func() error {
		log.Println("call stop callback")
		return nil
	})
	go app.MustRun()

	time.Sleep(time.Duration(3) * time.Second)
	app.Stop()
}
