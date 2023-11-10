package main

import (
	"fmt"
	"whoami/http"
	"whoami/uuid"
)

func main() {
	var server = http.Server{
		Address: "localhost",
		Port:    8080,
		Uuid:    uuid.GenUUID(),
	}

	err := server.Start()
	if err != nil {
		fmt.Println("Server encountered an error")
		panic(err)
	}

}
