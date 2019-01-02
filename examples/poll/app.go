package poll

import (
	"github.com/sirupsen/logrus"

	"github.com/reorx/gouken"
	pb "github.com/reorx/gouken/examples/poll/proto"
	"github.com/reorx/gouken/examples/poll/service"
)

type App struct {
	GApp *gouken.Application
}

func NewApp() *App {
	gapp := gouken.NewApplication(gouken.Config{
		Name:        "poll",
		Host:        "127.0.0.1",
		Port:        20001,
		Logger:      logrus.New(),
		LogRequest:  true,
		LogResponse: true,
	})
	gapp.SetDefaultInterceptor()

	app := &App{
		GApp: gapp,
	}

	return app
}

func (a *App) InitServer() {
	pb.RegisterPollServer(a.GApp.Server(), new(service.Poll))
}

func (a *App) GRPCClient() pb.PollClient {
	return pb.NewPollClient(a.GApp.Client())
}
