package test

import (
	"github.com/reorx/gouken/examples/poll"
	pb "github.com/reorx/gouken/examples/poll/proto"
)

var client pb.PollClient

// GetClient ...
func GetClient() pb.PollClient {
	return client
}

func init() {
	app := poll.NewApp()
	client = pb.NewPollClient(app.Client())
}
