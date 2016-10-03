package poll

import (
	"github.com/reorx/gouken"

	pb "github.com/reorx/gouken/examples/poll/proto"
	"github.com/reorx/gouken/examples/poll/service"
)

// NewApp ...
func NewApp() gouken.Application {
	gouken.MakeConfig(
		"config.yml",
		gouken.ConfPathEnv("POLLPATH"),
		gouken.ConfEnvPrefix("POLL"),
		gouken.ConfBindEnv("debug"),
	)

	app := gouken.NewApplication(
		gouken.Name("poll"),
	)

	// Register handler
	pb.RegisterPollServer(app.Server(), new(service.Poll))

	return app
}
