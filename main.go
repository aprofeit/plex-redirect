package main

import (
  "net/http"
  "os"
  "github.com/rs/zerolog/log"
)

func main() {
  port := os.Getenv("PORT")
  if port == "" {
    port = "8080"
  }

  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    http.Redirect(w, r, "https://forms.gle/xHfspeCYAvgTCyPp7", 307)
  })

  log.Info().Msgf("Starting server on port %s", port)
  log.Fatal().Err(http.ListenAndServe(":" + port, nil))
}
