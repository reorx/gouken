package gouken

import (
	"fmt"
	"log"
	"net"
	"strings"
	"sync"

	"golang.org/x/net/context"

	glog "github.com/reorx/gouken/log"
	"github.com/reorx/gouken/utils"
	"google.golang.org/grpc"
)

type application struct {
	// Options
	Name          string
	Host          string
	Port          int
	ClientAddress string
	LogLevel      string
	LogFilename   bool
	LogRequest    bool
	LogResponse   bool
	Debug         bool
	// Server
	server        *grpc.Server
	serverOnce    sync.Once
	Sopts         []grpc.ServerOption
	stopCallbacks []AppCallback
	listener      net.Listener
	// Client
}

var logResponse bool
var logRequest bool

func newApplication(opts ...Option) Application {
	// init with config
	a := &application{
		Name:          confName(),
		Host:          confHost(),
		Port:          confPort(),
		ClientAddress: confClientAddress(),
		LogLevel:      confLogLevel(),
		LogFilename:   confLogFilename(),
		LogRequest:    confLogRequest(),
		LogResponse:   confLogResponse(),
		Debug:         confDebug(),
		stopCallbacks: []AppCallback{},
	}

	// apply options
	for _, o := range opts {
		o(a)
	}

	// setup logging
	glog.Setup(a.LogLevel, a.LogFilename)

	// add interceptor
	a.Sopts = append(a.Sopts, grpc.UnaryInterceptor(applicationInterceptor))

	// print application
	log.Printf("%v created\n", a)
	return a
}

func (a *application) UseOptions(opts ...Option) {
	for _, o := range opts {
		o(a)
	}
}

func (a *application) Server() *grpc.Server {
	a.serverOnce.Do(func() {
		// set logRequest & logResponse
		logRequest = a.LogRequest
		logResponse = a.LogResponse

		// init server
		a.server = grpc.NewServer(a.Sopts...)
	})
	return a.server
}

func (a *application) Client() *grpc.ClientConn {
	addr := a.ClientAddress
	if addr == "" {
		addr = a.addr()
	}
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		glog.FatalKV("failed to connect", glog.Fields{"err": err, "address": addr})
	}
	log.Printf("client connected to %v", addr)
	return conn
}

func (a *application) Run() {
	a.listener = a.listen()
	a.server.Serve(a.listener)
}

type AppCallback func() error

func (a *application) OnStop(cb AppCallback) {
	a.stopCallbacks = append(a.stopCallbacks, cb)
}

func (a *application) Stop() {
	log.Printf("stop %v\n", a)
	a.listener.Close()
	for _, cb := range a.stopCallbacks {
		cb()
	}
}

func (a *application) PrintConfig() {
	PrintConfig()
}

func (a *application) String() string {
	return fmt.Sprintf("<application: Name=%v Host=%v Port=%v LogLevel=%v Debug=%v>",
		a.Name, a.Host, a.Port, a.LogLevel, a.Debug)
}

func (a *application) ServerAddress() string {
	return a.addr()
}

func (a *application) addr() string {
	return fmt.Sprintf("%v:%v", a.Host, a.Port)
}

func (a *application) listen() net.Listener {
	addr := a.addr()
	lis, err := net.Listen("tcp", a.addr())
	if err != nil {
		glog.FatalKV("failed to listen: %v", glog.Fields{"err": err, "address": addr})
	}
	log.Printf("server listening on %v\n", addr)
	return lis
}

func applicationInterceptor(ctx context.Context, req interface{},
	info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	s := strings.Split(info.FullMethod, "/")
	method := s[len(s)-1]
	// glog.Info("get request in interceptor 0 ", handler, req, ctx, info)
	// glog.Info("get request in interceptor 1 ", info)

	// call handler
	t0 := utils.NowTimestamp(13)

	resp, err = handler(ctx, req)

	tc0 := utils.NowTimestamp(13) - t0

	// log the call
	kvs := glog.Fields{
		"method": method,
		"ms":     tc0,
	}
	if logRequest {
		kvs["request"] = req
	}
	if logResponse {
		kvs["response"] = resp
	}
	logFunc := glog.InfoKV
	if err != nil {
		kvs["err"] = err
		logFunc = glog.ErrorKV
	}
	logFunc("/"+method+" called", kvs)

	return resp, err
}
