package main

import (
    "encoding/json"
    "net/http"
)

func mainPage(res http.ResponseWriter, req *http.Request) {
    res.Write([]byte("Привет!"))
}

func apiPage(res http.ResponseWriter, req *http.Request) {
    res.Write([]byte("Это страница /api."))
}

type Subj struct {
    Product string `json:"name"`
    Price   int    `json:"price"`
}

type jsonHandler struct {
    http.Handler
}

func (h *jsonHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    subj := Subj{"Milk", 50}
    resp, err := json.Marshal(subj)
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
    w.Header().Set("content-type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(resp)
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc(`/api/`, apiPage)
    mux.HandleFunc(`/`, mainPage)
    mux.Handle(`/json/`, &jsonHandler{})

    err := http.ListenAndServe(`:8082`, mux)
    if err != nil {
        panic(err)
    }
}