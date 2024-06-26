package main

import (
    "net/http"
)

func main() {
   // простейший сервер, которому доступны все файлы в поддиректории static
    err := http.ListenAndServe(":8088", http.FileServer(http.Dir("/static")))
    if err != nil {
        panic(err)
    }
}