package main

import (
  "net/http"
  )


func home(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("You found the Home Page"))
  }
