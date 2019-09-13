package main

import (
  "net/http"
  )


func (app *application) home(w http.ResponseWriter, r *http.Request) {
  //Changed this to render the home page template
  app.render(w, r, "home.page.tmpl", nil)
  }
  //copied the function to generate other pages
func (app *application) about(w http.ResponseWriter, r *http.Request) {
  app.render(w, r, "about.page.tmpl", nil)
  }
func (app *application) calendar(w http.ResponseWriter, r *http.Request) {
  app.render(w, r, "calendar.page.tmpl", nil)
  }
func (app *application) forum(w http.ResponseWriter, r *http.Request) {
  app.render(w, r, "forum.page.tmpl", nil)
  }
