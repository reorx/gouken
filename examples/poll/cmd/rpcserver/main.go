package main

import (
	"log"
	"time"

	pb "github.com/reorx/gouken/examples/poll/proto"
	"github.com/reorx/gouken/utils"

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
	go func() {
		c := app.GRPCClient()
		resp, err := c.AddCandidate(utils.Context(), &pb.AddCandidateRequest{
			Name: "foo",
		})
		log.Printf("got response=%s, err=%v", resp, err)
	}()

	time.Sleep(time.Duration(3) * time.Second)
	app.GApp.Stop()
}
