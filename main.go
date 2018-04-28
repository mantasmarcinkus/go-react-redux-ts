package main

import (
	"gopkg.in/macaron.v1"
	"log"
	"net/http"
	"algis/service/handler"
)

func main() {
	m := macaron.Classic()

	m.Use(macaron.Renderer(macaron.RenderOptions{
		// Directory to load templates. Default is "templates".
		Directory: "dist",
	}))

	m.Use(macaron.Static("dist"))

	m.Get("/", handler.IndexHandler)

	log.Println("Server is runnings...")
	log.Println("http://0.0.0.0:4000")
	http.ListenAndServe("0.0.0.0:4000", m)
	m.Run()
}
