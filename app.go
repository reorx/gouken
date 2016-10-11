package gouken

import (
	"fmt"
	"log"
	"net"
	"strings"

	"golang.org/x/net/context"

	glog "github.com/reorx/gouken/log"
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
	server *grpc.Server
	Sopts  []grpc.ServerOption
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
	if a.server == nil {
		// set logRequest & logResponse
		logRequest = a.LogRequest
		logResponse = a.LogResponse

		// init server
		a.server = grpc.NewServer(a.Sopts...)
	}
	return a.server
}

func (a *application) Client() *grpc.ClientConn {
	addr := a.ClientAddress
	if addr == "" {
		addr = a.addr()
	}
	conn, err := grpc.Dial(a.addr(), grpc.WithInsecure())
	if err != nil {
		glog.FatalKV("failed to connect", glog.Fields{"err": err, "address": addr})
	}
	log.Printf("client connected to %v", addr)
	return conn
}

func (a *application) Run() {
	a.run()
}

func (a *application) PrintConfig() {
	PrintConfig()
}

func (a *application) String() string {
	return fmt.Sprintf("<application: Name=%v Host=%v Port=%v LogLevel=%v Debug=%v>",
		a.Name, a.Host, a.Port, a.LogLevel, a.Debug)
}

func (a *application) addr() string {
	return fmt.Sprintf("%v:%v", a.Host, a.Port)
}

func (a *application) run() {
	addr := a.addr()
	lis, err := net.Listen("tcp", a.addr())
	if err != nil {
		glog.FatalKV("failed to listen: %v", glog.Fields{"err": err, "address": addr})
	}
	log.Printf("server listening on %v\n", addr)

	a.server.Serve(lis)
}

func applicationInterceptor(ctx context.Context, req interface{},
	info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	s := strings.Split(info.FullMethod, "/")
	method := s[len(s)-1]
	// glog.Info("get request in interceptor 0 ", handler, req, ctx, info)
	// glog.Info("get request in interceptor 1 ", ctx)
	// glog.Info("get request in interceptor 2 ", info)
	reqf := glog.Fields{"method": method}
	if logRequest {
		reqf["request"] = req
	}
	glog.InfoKV("/"+method+" received", reqf)

	resp, err = handler(ctx, req)

	// log err
	resps := "/" + method + " responded"
	if err != nil {
		respf := glog.Fields{"err": err, "method": method}
		if logResponse {
			respf["response"] = resp
		}
		glog.ErrorKV(resps+" with error", respf)
	} else {
		if logResponse {
			glog.InfoKV(resps, glog.Fields{"response": resp, "method": method})
		}
	}

	return resp, err
}
