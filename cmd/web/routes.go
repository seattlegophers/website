package main

import (
  "net/http"
  "github.com/bmizerany/pat"
  "github.com/justinas/alice"
  )


func (app *application) routes() http.Handler {

  dynamicMiddleware := alice.New(app.session.Enable)

  mux := pat.New()
//added other routes
  mux.Get("/", dynamicMiddleware.ThenFunc(app.home))
  mux.Get("/about", dynamicMiddleware.ThenFunc(app.about))
  mux.Get("/calendar", dynamicMiddleware.ThenFunc(app.calendar))
  mux.Get("/forum", dynamicMiddleware.ThenFunc(app.forum))
  mux.Get("/user/signin", dynamicMiddleware.ThenFunc(app.signinForm))
  mux.Post("/user/signup", dynamicMiddleware.ThenFunc(app.signUp))//Implement this


  return mux
  }


