package main

import (
	"flag"

	"goazl/config"
	"goazl/server"
)

func main() {
	port := flag.String("port", "8989", "Specify the port to run the server.")
	env := flag.String("env", "release", "Switch between release or debug mode.")

	flag.Parse()

	config := config.Config{
		Environment: *env,
		Port:        *port,
	}

	server.Run(&config)
}
