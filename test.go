package main

import (
    "fmt"
    "net/http"
    "os"
)

func main() {
    http.HandleFunc("/", gogogo)
    fmt.Println("listening...")
    err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
    
    if err != nil {
      panic(err)
    }
}

func gogogo(res http.ResponseWriter, req *http.Request) {
    fmt.Fprintln(res, "go johnny go go go go")
}
