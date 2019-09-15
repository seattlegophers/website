package main

import (
  "net/http"
  "github.com/bmizerany/pat"
  "github.com/justinas/alice"
  )


func (app *application) routes() http.Handler {

  dynamicMiddleware := alice.New(app.session.Enable, app.authenticate)

  mux := pat.New()
//added other routes
  mux.Get("/", dynamicMiddleware.ThenFunc(app.home))
  mux.Get("/about", http.HandlerFunc(app.about))
  mux.Get("/calendar", http.HandlerFunc(app.calendar))
  mux.Get("/forum", http.HandlerFunc(app.forum))
  mux.Get("/user/signin", http.HandlerFunc(app.signinForm))
  mux.Post("/user/signup", http.HandlerFunc(app.signUp))//Implement this


  return mux
  }


