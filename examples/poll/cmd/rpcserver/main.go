package main

import (
	"log"
	"time"

	"github.com/reorx/gouken/examples/poll"
)

func main() {
	app := poll.NewApp()
	app.InitServer()

	app.GApp.OnStop(func() error {
		log.Println("call stop callback")
		return nil
	})
	go app.GApp.MustRun()

	time.Sleep(time.Duration(3) * time.Second)
	app.GApp.Stop()
}
