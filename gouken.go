package gouken

import "google.golang.org/grpc"

// Application is an interface for building and initialising application.
type Application interface {
	Run()
	Configure(opts ...Option)
	Server() *grpc.Server
	Client() *grpc.ClientConn
	String() string
}

// NewApplication creates an returns a new Application based on the packages within.
func NewApplication(opts ...Option) Application {
	return newApplication(opts...)
}
