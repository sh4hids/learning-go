package main

import (
  "fmt"
  "net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "Hello there from goland...!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "This is the about page...")
}

func main() {
  http.HandleFunc("/", indexHandler)
  http.HandleFunc("/about", aboutHandler)
  http.ListenAndServe(":8000", nil)
}
