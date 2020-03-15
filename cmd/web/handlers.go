package main

import (
  "fmt"
  "net/http"
  "strconv"
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
  // Extract the value of the id parameter from the query string and try to
  // convert it to an integer using the strconv.Atoi() function. If it can't
  // be converted to an integer, or the value is less than 1, we return a 404
  // page not found response.
  id,err := strconv.Atoi(r.URL.Query().Get("id"))
  if err != nil || id < 1 {
    http.NotFound(w,r)
    return
  }

  // Use the fmt.Fprintf() function to interpolate the id value with our response
  // and write it to the http.ResponseWriter.
  fmt.Fprintf(w, "Display a specific snippet with ID %d...\n\n", id)
  if id % 2 == 0 {
    w.Write([]byte("Additionally, it was divisible by 2!"))
  }
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
    http.Error(w, "Method not allowed", 405)
    return
  }
  w.Header().Set("Content-Type", "application/json")
  w.Write([]byte(`{"name":"Alex"}`))
  //w.Write([]byte("Creating a snippet"))
}
