package main

import (
  "log"
  "net/http"
)

func main() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    http.Redirect(w, r, "https://forms.gle/xHfspeCYAvgTCyPp7", 307)
  })

  log.Fatal(http.ListenAndServe(":8080", nil))
}
