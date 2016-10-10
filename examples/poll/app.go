package poll

import (
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
	gouken.MakeConfig(
		"config.yml",
		gouken.ConfPathEnv("POLLPATH"),
		gouken.ConfEnvPrefix("POLL"),
		gouken.ConfBindEnv("debug"),
	)

	app = gouken.NewApplication(
		gouken.Name("poll"),
	)

	// Register handler
	pb.RegisterPollServer(app.Server(), new(service.Poll))
}
