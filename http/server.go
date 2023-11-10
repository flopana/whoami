package http

import (
	"fmt"
	"net/http"
)

type Server struct {
	Address string
	Port    uint16
	Uuid    string
}

func (s *Server) Start() error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprint(w, "Welcome to my website!")
		if err != nil {
			fmt.Println("Error writing to response writer")
			panic(err)
		}
	})

	var fullAddress = fmt.Sprintf("%s:%d", s.Address, s.Port)
	fmt.Println("Server started on address: " + fullAddress)

	return http.ListenAndServe(fullAddress, nil)
}
