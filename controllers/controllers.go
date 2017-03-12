package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("views/*"))
}

// Index home page
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func checkErr(wr http.ResponseWriter, err error) {
	if err != nil {
		errorResponse(wr, err.Error(), http.StatusInternalServerError)
		return
	}
}

// jsonResponse write response to the user
func jsonResponse(wr http.ResponseWriter, objJSON interface{}, sttsCode int) {
	wr.Header().Set("Content-Type", "application/json; charset=UTF-8")
	wr.WriteHeader(sttsCode)
	fmt.Fprintf(wr, "%s\n", objJSON)
}

func errorResponse(wr http.ResponseWriter, msg string, sttsCode int) {
	wr.Header().Set("Content-Type", "application/json; charset=UTF-8")
	wr.WriteHeader(sttsCode)
	fmt.Fprintf(wr, "{message: %s}", msg)
}
