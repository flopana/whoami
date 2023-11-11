package main

import (
	"fmt"
	"image/color"
	"math/rand"
	"whoami/http"
	"whoami/uuid"
)

func main() {
	var server = http.Server{
		Address: "0.0.0.0",
		Port:    8080,
		Uuid:    uuid.GenUUID(),
		Color:   getColor(),
	}

	err := server.Start()
	if err != nil {
		fmt.Println("Server encountered an error")
		panic(err)
	}

}

func getColor() color.RGBA {
	return color.RGBA{
		R: uint8(rand.Intn(256)),
		G: uint8(rand.Intn(256)),
		B: uint8(rand.Intn(256)),
		A: 255,
	}
}
