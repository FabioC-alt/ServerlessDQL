package main

import (
	"fmt"
	
	"os"
    "time"
)

func main() {

	appEnv := os.Getenv("APP_ENV")
	Ports := os.Getenv("PORTS")

	i := 1
	for {
		fmt.Printf("Loop %d - App Environment: %s\n", i, appEnv)
		fmt.Printf("Loop %d - Ports: %s\n", i, Ports)
		i++
		time.Sleep(2 * time.Second)
	}
}