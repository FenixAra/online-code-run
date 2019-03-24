package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/FenixAra/online-code-run/dtos"
	"github.com/FenixAra/online-code-run/internal/services/api"

	"github.com/julienschmidt/httprouter"
)

func GetLanguages(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Add("X-Content-Type-Options", "nosniff")
	w.Header().Add("X-Frame-Options", "DENY")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	languages := api.GetLanguages()

	response, err := json.Marshal(languages)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func CompileCode(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

}

func CompileAndRunCode(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Add("X-Content-Type-Options", "nosniff")
	w.Header().Add("X-Frame-Options", "DENY")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	decoder := json.NewDecoder(r.Body)
	req := &dtos.APIReq{}

	err := decoder.Decode(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	res := api.CompileAndRunCode(req)
	response, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
