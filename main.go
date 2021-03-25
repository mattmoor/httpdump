package main

import (
  "fmt"
  "log"
  "net/http"
  "net/http/httputil"
  "os"
)

func handler(w http.ResponseWriter, r *http.Request) {
  b, err := httputil.DumpRequest(r, true /* dump body */)
  if err != nil {
	http.Error(w, err.Error(), http.StatusInternalServerError)
  }

  w.Write([]byte(fmt.Sprintf("GOT: %s", string(b))))
}

func main() {
  http.HandleFunc("/", handler)

  port := os.Getenv("PORT")
  if port == "" {
    port = "8080"
  }
  log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}