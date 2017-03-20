package main

import (
    "io"
    "log"
    "net/http"
)

func helloHandle(w http.ResponseWriter,r *http.Request){
    io.WriteString(w, "Hello World!")
}

func main(){
    http.HandleFunc("/hello", helloHandle)
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err.Error())
    }
}
