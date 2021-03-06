package main

import (
  "log"
  "net/http"
)



func main(){
  //Use the http.NewServeMux() function to init a new servemux, then
  // register the home function as the handler for the '/' route

  mux := http.NewServeMux()
  mux.HandleFunc("/", home)
  mux.HandleFunc("/snippet", showSnippet)
  mux.HandleFunc("/snippet/create", createSnippet)

  // Use http.ListenAndServe() to start webserver. Takes in two params
  // port and the mux object. If it returns an error we progress in
  // the program to log.Fatal() function to log the error and exit.
  // Note: Any error returned by http.ListenAndServe() is always non-nil

  log.Println("starting server on port 4000")
  err := http.ListenAndServe("localhost:4000", mux)
  log.Fatal(err)
}
