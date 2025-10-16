package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	appEnv := os.Environ("APP_ENV")
	Address := os.Environ("ADDRESS")
	Ports := os.Environ("PORTS")

	Printf("App Environment: %s", appEnv)
	Printf("Address: %s", Address)
	Printf("Ports: %s", Ports)
	

}