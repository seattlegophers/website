package main

import (
  "net/http"
  "github.com/bmizerany/pat" //allows us to do request based routing
  )


func routes() http.Handler {

  mux := pat.New()

  mux.Get("/", http.HandlerFunc(home))


  return mux
  }


