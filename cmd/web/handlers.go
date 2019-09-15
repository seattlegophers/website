package main

import (
  "net/http"
  "seattleGophers.com/website/pkg/forms"
)


func (app *application) home(w http.ResponseWriter, r *http.Request) {
  app.render(w, r, "home.page.tmpl", nil)
}
func (app *application) about(w http.ResponseWriter, r *http.Request) {
  app.render(w, r, "about.page.tmpl", nil)
}
func (app *application) calendar(w http.ResponseWriter, r *http.Request) {
  app.render(w, r, "calendar.page.tmpl", nil)
}
func (app *application) forum(w http.ResponseWriter, r *http.Request) {
  app.render(w, r, "forum.page.tmpl", nil)
}
func (app *application) signinForm(w http.ResponseWriter, r *http.Request) {
  app.render(w, r, "signin.page.tmpl", &templateData{
    Form:  forms.New(nil),
  })
}
func (app *application) signUp(w http.ResponseWriter, r *http.Request) {
  app.render(w, r, "home.page.tmpl", nil)//placeholder
}
