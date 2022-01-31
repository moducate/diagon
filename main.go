package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/moducate/x/errorsx"
	"github.com/moducate/x/osx"
)

type WondeResponse struct {
	Data []interface{} `json:"data"`
	Meta struct {
		Pagination struct {
			Next string `json:"next"`
		} `json:"pagination"`
	} `json:"meta"`
}

func handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		errorsx.New(http.StatusMethodNotAllowed, "method-not-allowed", "HTTP Method Not Allowed", r.URL.Path, nil).WriteToHttp(w)
		return
	}

	schools := strings.Split(r.Header.Get("x-schools"), ",")

	if len(schools) == 1 && schools[0] == "" {
		errorsx.New(http.StatusBadRequest, "missing-header", "Missing HTTP Request Header", r.URL.Path, map[string]interface{}{"header": "x-schools"}).WriteToHttp(w)
		return
	}

	for _, school := range schools {
		if strings.TrimSpace(school) == "" {
			continue
		}

		next := fmt.Sprintf("https://api.wonde.com/v1.0/schools/%s%s", school, r.URL.Path)
		if r.URL.RawQuery != "" {
			next += fmt.Sprintf("?%s", r.URL.RawQuery)
		}

		for next != "" {

		}
	}
}

func main() {
	http.HandleFunc("/", handle)
	http.ListenAndServe(osx.Getenv("ADDR", ":3000"), nil)
}
