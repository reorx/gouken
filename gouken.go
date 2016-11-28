package gouken

import "google.golang.org/grpc"

// Application is an interface for building and initialising application.
type Application interface {
	Run()
	OnStop(AppCallback)
	Stop()
	UseOptions(opts ...Option)
	Server() *grpc.Server
	ServerAddress() string
	Client() *grpc.ClientConn
	PrintConfig()
	String() string
}

// NewApplication creates an returns a new Application based on the packages within.
func NewApplication(opts ...Option) Application {
	return newApplication(opts...)
}

func init() {
	defineConfig()
}
