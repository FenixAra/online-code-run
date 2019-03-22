package main

import (
	"log"
	"net/http"

	"github.com/FenixAra/online-code-run/internal/config"
	"github.com/FenixAra/online-code-run/internal/handlers"
)

func main() {
	log.Println("Listening to port: ", config.Port)
	http.ListenAndServe(":"+config.Port, handlers.New())
}
