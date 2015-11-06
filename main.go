package main

import (
	"net/http"
  "fmt"
  "github.com/toqueteos/webbrowser"
)

func main() {
   http.Handle("/", http.FileServer(assetFS()))
   fmt.Println("Serving on port 8000")
	 webbrowser.Open("http://localhost:8000")
   http.ListenAndServe(":8000", nil)
}
