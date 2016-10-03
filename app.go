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
	Name     string
	Host     string
	Port     int
	LogLevel string
	Debug    bool
	// Server
	server *grpc.Server
	Sopts  []grpc.ServerOption
	// Client
}

func newApplication(opts ...Option) Application {
	// init with config
	a := &application{
		Name:     confName(),
		Host:     confHost(),
		Port:     confPort(),
		LogLevel: confLogLevel(),
		Debug:    confDebug(),
	}

	// apply options
	for _, o := range opts {
		o(a)
	}

	// setup logging
	glog.Setup(a.LogLevel)

	// add interceptor
	a.Sopts = append(a.Sopts, grpc.UnaryInterceptor(applicationInterceptor))

	// init server
	a.server = grpc.NewServer(a.Sopts...)

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
	return a.server
}

func (a *application) Client() *grpc.ClientConn {
	conn, err := grpc.Dial(a.addr(), grpc.WithInsecure())
	if err != nil {
		glog.FatalKV("failed to connect", glog.Fields{"err": err})
	}
	return conn
}

func (a *application) Run() {
	a.run()
}

func (a *application) String() string {
	return fmt.Sprintf("<application: Name=%v Host=%v Port=%v LogLevel=%v Debug=%v>",
		a.Name, a.Host, a.Port, a.LogLevel, a.Debug)
}

func (a *application) addr() string {
	return fmt.Sprintf("%v:%v", a.Host, a.Port)
}

func (a *application) run() {
	lis, err := net.Listen("tcp", a.addr())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("listening on %v\n", lis.Addr())

	a.server.Serve(lis)
}

func applicationInterceptor(ctx context.Context, req interface{},
	info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	s := strings.Split(info.FullMethod, "/")
	method := s[len(s)-1]
	// glog.Info("get request in interceptor 0 ", handler, req, ctx, info)
	// glog.Info("get request in interceptor 1 ", ctx)
	// glog.Info("get request in interceptor 2 ", info)
	glog.InfoKV("/"+method+" start", glog.Fields{"req": req, "method": method})

	resp, err = handler(ctx, req)

	// log err

	glog.InfoKV("/"+method+" end", glog.Fields{"resp": resp, "method": method})
	return resp, err
}
