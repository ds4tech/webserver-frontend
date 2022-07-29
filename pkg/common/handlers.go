package common

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var dat map[string]interface{}

func HomeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to simple Webserver!")
}

func Health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Here the logs come:")
}

func EchoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	txt := r.URL.Query().Get("text")
	if len([]rune(txt)) < 1 {
		txt = "Text not provided in a query string."
	}

	jsonMap := map[string]string{"result": txt}
	jsonResult, _ := json.Marshal(jsonMap)
	fmt.Fprintf(w, string(jsonResult))
}
