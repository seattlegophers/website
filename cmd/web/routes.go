package main

import (
  "net/http"
  "github.com/bmizerany/pat" //allows us to do request based routing
  )


func (app *application) routes() http.Handler {

  mux := pat.New()
//added other routes
  mux.Get("/", http.HandlerFunc(app.home))
  mux.Get("/about", http.HandlerFunc(app.about))
  mux.Get("/calendar", http.HandlerFunc(app.calendar))
  mux.Get("/forum", http.HandlerFunc(app.forum))
  mux.Get("/user/signin", http.HandlerFunc(app.signinForm))


  return mux
  }


