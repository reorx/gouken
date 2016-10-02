package poll

import (
	"fmt"

	"github.com/reorx/gouken"

	pb "github.com/reorx/gouken/examples/poll/proto"
	"github.com/reorx/gouken/examples/poll/service"
)

// NewApp ...
func NewApp() gouken.Application {
	InitConfig()

	app := gouken.NewApplication(
		gouken.Name("poll"),
		gouken.Port(9090),
		gouken.ConfLogLevel("log_level"),
	)

	fmt.Println(app)

	// Register handler
	pb.RegisterPollServer(app.Server(), new(service.Poll))

	return app
}
