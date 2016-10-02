package test

import (
	"github.com/reorx/gouken/examples/poll"
	pb "github.com/reorx/gouken/examples/poll/proto"
)

// NewTestClient ...
func NewTestClient() pb.PollClient {
	app := poll.NewApp()
	return pb.NewPollClient(app.Client())
}
