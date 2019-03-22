package handlers

import (
	"log"
	"net/http"
	"runtime/debug"

	"github.com/julienschmidt/httprouter"
)

func New() http.Handler {
	router := httprouter.New()
	router.PanicHandler = PanicHandler

	router.GET("/api/v1/languages", GetLanguages)

	return router
}

func PanicHandler(w http.ResponseWriter, r *http.Request, c interface{}) {
	log.Printf("Recovering from panic, Reason: %+v", c.(error))
	debug.PrintStack()
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(c.(error).Error()))
}
