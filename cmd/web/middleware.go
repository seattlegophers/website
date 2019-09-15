package main

import (
  "context"
  "net/http"
  "seattleGophers.com/website/pkg/models"
)

func (app *application) authenticate(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    exists := app.session.Exists(r, "authenticatedUserID")
    if !exists {
      next.ServeHTTP(w, r)
      return
    }

    user, err := app.users.Get(app.session.GetInt(r, "authenticatedUserID"))
    if err == models.ErrNoRecord || !user.Active {
      app.session.Remove(r, "authenticatedUserID")
      next.ServeHTTP(w, r)
      return
    } else if err != nil {
      app.serverError(w, err)
      return
    }

    ctx := context.WithValue(r.Context(), contextKeyIsAuthenticated, true)
    next.ServeHTTP(w, r.WithContext(ctx))
  })
}

