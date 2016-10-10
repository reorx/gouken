package main

import (
	"github.com/reorx/gouken/examples/poll"
)

func main() {
	app := poll.App()
	app.PrintConfig()
	app.Run()
}
