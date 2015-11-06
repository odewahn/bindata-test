package main

import (
	"net/http"
  "fmt"
)

func main() {
   http.Handle("/", http.FileServer(assetFS()))
   fmt.Println("Serving on port 8000")
   http.ListenAndServe(":8000", nil)
}
