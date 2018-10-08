package poll

import (
	"github.com/sirupsen/logrus"

	"github.com/reorx/gouken"
	pb "github.com/reorx/gouken/examples/poll/proto"
	"github.com/reorx/gouken/examples/poll/service"
)

var app gouken.Application
var client pb.PollClient

// App ...
func App() gouken.Application {
	Init()
	return app
}

// Client ..
func Client() pb.PollClient {
	if client == nil {
		client = pb.NewPollClient(app.Client())
	}
	return client
}

// Init ..
func Init() {
	app = gouken.NewApplication(gouken.Config{
		Name:        "poll",
		Host:        "127.0.0.1",
		Port:        20001,
		Logger:      logrus.New(),
		LogRequest:  true,
		LogResponse: true,
	})

	// Register handler
	pb.RegisterPollServer(app.Server(), new(service.Poll))
}
