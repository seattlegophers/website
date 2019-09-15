package main

import (
  "flag"
  "log"
  "crypto/tls"
  "os"
  "net/http"
  "html/template"
  "github.com/golangcollege/sessions"
  "seattleGophers.com/website/pkg/models"
  )

type contextKey string
var contextKeyIsAuthenticated = contextKey("isAuthenticated")

type application struct {
  infoLog          *log.Logger
  errorLog         *log.Logger
  templateCache    map[string]*template.Template
  session          *sessions.Session
  users            interface {
    Insert(string, string, string) error
    Authenticate(string, string) (int, error)
    Get(int) (*models.User, error)
  }
}




func main() {
  addr := flag.String("addr", ":8080", "Port to accept incoming connections")
  flag.Parse()

  infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
  errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

  templateCache, err := newTemplateCache("./ui/html/") // Puts template files into cache for easy access
  if err != nil {
    errorLog.Fatal(err)
    }

  app := &application {
    errorLog:        errorLog,
    infoLog:         infoLog,
    templateCache:   templateCache,
    }

  tlsConfig := &tls.Config {
    PreferServerCipherSuites: true,
    CurvePreferences:         []tls.CurveID{tls.X25519, tls.CurveP256},
    }
  srv := &http.Server {
    Addr:         *addr,
    ErrorLog:     errorLog,
    Handler:      app.routes(), // see cmd/web/routes.go
    TLSConfig:    tlsConfig,
    }
  infoLog.Printf("Starting server on %s", *addr)
  //   $ cd tls; go run /usr/local/go/src/crypto/tls/generate_cert.go --rsa-bits=2048 --host=localhost
  err = srv.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem")
  errorLog.Fatal(err)
  }
