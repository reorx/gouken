package gouken

import "google.golang.org/grpc"

var _ Application = new(application)

// Application is an interface for building and initialising application.
type Application interface {
	MustRun()
	OnStop(AppCallback)
	Stop()
	Server() *grpc.Server
	Client() *grpc.ClientConn
}

// NewApplication creates an returns a new Application based on the packages within.
func NewApplication(config Config) Application {
	return newApplication(config)
}
