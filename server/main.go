package main

import (
	"fmt"
	"net/http"
)

func main () {
    err := http.ListenAndServe(":9091", http.FileServer(http.Dir("../examples")))
    if err != nil {
        fmt.Println("Failed to start server", err)
        return
    }
}
