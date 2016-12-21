package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/technoweenie/grohl"
)

const (
	CONTENT_TYPE_PROBLEM_JSON = "application/problem+json"
)

type Problem struct {
	Type   string `json:"type,omitempty"`
	Title  string `json:"title,omitempty"`
	Status int    `json:"status,omitempty"`
	Detail string `json:"detail,omitempty"`
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	grohl.Log(grohl.Data{"method": r.Method, "path": r.URL.Path, "status": "404"})
	Error(w, r, Problem{Status: http.StatusNotFound}, http.StatusNotFound)
}

func Error(w http.ResponseWriter, r *http.Request, problem Problem, status int) {
	w.Header().Set("Content-Type", CONTENT_TYPE_PROBLEM_JSON)
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(status)
	problem.Status = status

	jsonError, e := json.Marshal(problem)
	if e != nil {
		http.Error(w, "{}", problem.Status)
	}
	fmt.Fprintln(w, string(jsonError))
}
