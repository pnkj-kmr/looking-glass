package main

import (
	"flag"
	"fmt"

	"github.com/pnkj-kmr/looking-glass/handlers"
	"github.com/pnkj-kmr/looking-glass/utils"
)

func main() {
	var port string
	var debug bool
	flag.StringVar(&port, "port", "8080", "Application Port")
	flag.BoolVar(&debug, "debug", false, "Application Debug Mode")
	flag.Parse()

	// Setting up the logger
	utils.SetLogger(debug)
	// Creating new server variable
	server := handlers.NewServer(debug, utils.L)

	// Starting application
	server.Run(fmt.Sprintf("0.0.0.0:%s", port))
}
