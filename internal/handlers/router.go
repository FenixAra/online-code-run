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
	router.POST("/api/v1/run", CompileAndRunCode)
	router.GET("/ping", Ping)

	return router
}

func Ping(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte("pong"))
}

func PanicHandler(w http.ResponseWriter, r *http.Request, c interface{}) {
	log.Printf("Recovering from panic, Reason: %+v", c.(error))
	debug.PrintStack()
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(c.(error).Error()))
}
