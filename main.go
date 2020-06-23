package main

import (
  "net/http"
  "os"
)

func main() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    http.Redirect(w, r, "https://forms.gle/xHfspeCYAvgTCyPp7", 307)
  })

  port := os.Getenv("PORT")
  if port == "" {
    port = "8080"
  }

  http.ListenAndServe(":" + port, nil)
}
