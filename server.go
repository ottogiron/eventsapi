package main

import (
	"fmt"
	"os"

	"github.com/ottogiron/chapi/server"
	"github.com/ottogiron/eventsapi/controllers"
)

func main() {
	connectionSTring := ":" + os.Getenv("PORT")
	server := server.NewServer()
	server.Register(&controllers.HelloController{})
	runError := server.Run(connectionSTring)
	if runError != nil {
		fmt.Println("Error when running server", runError)
	}
}
