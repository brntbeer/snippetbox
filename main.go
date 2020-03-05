package main

import (
  "log"
  "net/http"
)

// Define a home handler function which writes a byte slice containing
// "hello from snippetbox" at the response body
func home(w http.ResponseWriter, r *http.Request) {

  // 404 if the path isn't actually `/`
  if r.URL.Path != "/" {
    http.NotFound(w,r)
    return
  }

  w.Write([]byte("Hello from Snippetbox"))
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Showing Snippets"))
}

func createSnippet(w http.ResponseWriter, r *http.Request) {

  // Use r.Method to check whether the request is using Post or not
  // http.MethodPost is a constant equal to string "POST"
  if r.Method != http.MethodPost {
    // If it's not, use the w.WriteHeader() method to send a 405 status
    // code and the w.Write() method to write a "Method Not Allowed"
    // response body. We then return from the function so that the
    // subsequent code is not executed.
    w.Header().Set("Allow", http.MethodPost)
    w.WriteHeader(405)
    w.Write([]byte("Method Not Allowed"))
    return
  }
  w.Write([]byte("Creating a snippet"))
}

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
  err := http.ListenAndServe(":4000", mux)
  log.Fatal(err)
}
