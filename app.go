package gouken

import (
	"context"
	"fmt"
	"net"
	"strings"
	"sync"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"github.com/reorx/gouken/utils"
)

type Application struct {
	config Config
	// Server
	server        *grpc.Server
	serverOnce    sync.Once
	serverOpts    []grpc.ServerOption
	stopCallbacks []AppCallback
	listener      net.Listener
}

func NewApplication(config Config) *Application {
	if err := config.Check(); err != nil {
		panic(err)
	}
	// init with config
	a := &Application{
		config:        config,
		stopCallbacks: []AppCallback{},
	}

	// print application
	a.config.Logger.Infof("%v created", a)
	return a
}

func (a *Application) AppendServerOptions(serverOpts ...grpc.ServerOption) {
	a.serverOpts = append(a.serverOpts, serverOpts...)
}

func (a *Application) Server() *grpc.Server {
	a.serverOnce.Do(func() {
		// init server
		a.server = grpc.NewServer(a.serverOpts...)
	})
	return a.server
}

func (a *Application) Client() *grpc.ClientConn {
	addr := a.config.addr()
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		a.config.Logger.WithError(err).WithField("address", addr).Panic("failed to connect")
	}
	a.config.Logger.Infof("client connected to %v", addr)
	return conn
}

func (a *Application) MustRun() {
	a.listener = a.listen()
	a.server.Serve(a.listener)
}

type AppCallback func() error

func (a *Application) OnStop(cb AppCallback) {
	a.stopCallbacks = append(a.stopCallbacks, cb)
}

func (a *Application) Stop() {
	a.config.Logger.Infof("stop %v", a)
	a.listener.Close()
	for _, cb := range a.stopCallbacks {
		cb()
	}
}

func (a *Application) String() string {
	return fmt.Sprintf("<application: Name=%v Host=%v Port=%v Debug=%v>",
		a.config.Name, a.config.Host, a.config.Port, a.config.Debug,
	)
}

func (a *Application) listen() net.Listener {
	addr := a.config.addr()
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		a.config.Logger.WithError(err).WithField("address", addr).Panic("failed to listen on address")
	}
	a.config.Logger.Infof("server listening on %v", addr)
	return lis
}

func (a *Application) SetDefaultInterceptor() {
	a.AppendServerOptions(grpc.UnaryInterceptor(a.LoggingInterceptor()))
}

func (a *Application) LoggingInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		s := strings.Split(info.FullMethod, "/")
		method := s[len(s)-1]
		// glog.Info("get request in interceptor 0 ", handler, req, ctx, info)
		// glog.Info("get request in interceptor 1 ", info)

		// call handler
		t0 := utils.NowTimestamp(13)

		resp, err = handler(ctx, req)

		tc0 := utils.NowTimestamp(13) - t0

		// log the call
		kvs := logrus.Fields{
			"method": method,
			"ms":     tc0,
		}
		if a.config.LogRequest {
			kvs["request"] = req
		}
		if a.config.LogResponse {
			kvs["response"] = resp
		}

		message := "/" + method + " called"
		if err != nil {
			a.config.Logger.WithError(err).WithFields(kvs).Error(message)
		} else {
			a.config.Logger.WithFields(kvs).Info(message)
		}
		return resp, err
	}
}
