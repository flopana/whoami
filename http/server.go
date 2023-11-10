package http

import (
	"fmt"
	"html/template"
	"image/color"
	"net/http"
	"os"
	"strings"
)

type Server struct {
	Address string
	Port    uint16
	Uuid    string
	Color   color.RGBA
}

func (s *Server) Start() error {
	http.HandleFunc("/", s.servePage)

	replaceColorInTemplate(s)
	var fullAddress = fmt.Sprintf("%s:%d", s.Address, s.Port)
	fmt.Println("Server started on address: " + fullAddress)

	return http.ListenAndServe(fullAddress, nil)
}

func replaceColorInTemplate(s *Server) {
	file, err := os.Open("assets/index.html")
	defer file.Close()
	if err != nil {
		return
	}

	fileInfo, _ := file.Stat()

	var fileContents = make([]byte, fileInfo.Size())
	_, err = file.Read(fileContents)
	if err != nil {
		return
	}

	var fileContentsString = string(fileContents)
	var colors = fmt.Sprintf("rgba(%d, %d, %d, %d)", s.Color.R, s.Color.G, s.Color.B, s.Color.A)

	fileContentsString = strings.Replace(fileContentsString, "#!#identityColor#!#", colors, 1)

	err = os.WriteFile("assets/index.tmp.html", []byte(fileContentsString), 0644)
	if err != nil {
		return
	}
}

func (s *Server) servePage(w http.ResponseWriter, r *http.Request) {
	var pageData = Generate(s)

	var tmpl = template.Must(template.ParseFiles("assets/index.tmp.html"))

	err := tmpl.Execute(w, pageData)
	if err != nil {
		fmt.Println("Error executing template")
	}
}
