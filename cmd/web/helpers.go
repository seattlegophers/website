package main

import (
  "net/http"
  "fmt"
  "bytes"
  "runtime/debug"
  "time"
)

func (app *application) render(w http.ResponseWriter, r *http.Request, name string, td *templateData) {
  ts, ok := app.templateCache[name]
  if !ok {
    app.serverError(w, fmt.Errorf("the template %s does not exist", name))
    return
  }

  buffer := new(bytes.Buffer) //buffer the page before write

  err := ts.Execute(buffer, app.addDefaultData(td, r))
  if err != nil {
    app.serverError(w, err)
    return
  }
  buffer.WriteTo(w)
}

//Wrote error logic for writing errors to users
func (app *application) serverError(w http.ResponseWriter, err error) {
  trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
  app.errorLog.Output(2, trace)
  http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
func (app *application) clientError(w http.ResponseWriter, status int) {
  http.Error(w, http.StatusText(status), status)
}
func (app *application) isAuthenticated(r *http.Request) bool {
  return app.session.Exists(r, "authenticatedUserID")
}

func (app *application) addDefaultData(td *templateData, r *http.Request) *templateData {
  if td == nil {
    td = &templateData{}
  }

  td.CurrentYear = time.Now().Year()

  td.IsAuthenticated = app.isAuthenticated(r)
  return td
}
