package main

import (
	"flag"
	"log"
	"time"

	pb "github.com/reorx/gouken/examples/poll/proto"
	"github.com/reorx/gouken/utils"

	"github.com/reorx/gouken/examples/poll"
)

func main() {
	var oneOff bool
	flag.BoolVar(&oneOff, "one-off", true, "run the service, call a method, then stop and exit.")

	app := poll.NewApp()
	app.InitServer()

	app.GApp.OnStop(func() error {
		log.Println("call stop callback")
		return nil
	})

	go func() {
		c := app.GRPCClient()
		resp, err := c.AddCandidate(utils.Context(), &pb.AddCandidateRequest{
			Name: "foo",
		})
		log.Printf("got response=%s, err=%v", resp, err)
	}()

	if oneOff {
		go func() {
			time.Sleep(time.Duration(3) * time.Second)
			log.Println("call stop")
			app.GApp.Stop()
		}()
	}

	app.GApp.MustRun()
}
