package main

import (
  "net/http"
  "github.com/rs/zerolog/log"
)

func main() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    http.Redirect(w, r, "https://forms.gle/xHfspeCYAvgTCyPp7", 307)
  })

  log.Info().Msg("Starting server")
  log.Fatal().Err(http.ListenAndServe(":8080", nil))
}
