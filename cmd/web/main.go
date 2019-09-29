package main

import (
  "time"
  "flag"
  "log"
  "crypto/tls"
  "os"
  "net/http"
  "html/template"
  "github.com/golangcollege/sessions"
  "seattleGophers.com/website/pkg/models"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
  )

type contextKey string
var contextKeyIsAuthenticated = contextKey("isAuthenticated")

type application struct {
  infoLog          *log.Logger
  errorLog         *log.Logger
  templateCache    map[string]*template.Template
  session          *sessions.Session
  IsAuthenticated  bool
  users            interface {
    Insert(string, string, string) error
    Authenticate(string, string) (int, error)
    Get(int) (*models.User, error)
  }
}

func main() {
  addr := flag.String("addr", ":8080", "Port to accept incoming connections")
  secret := flag.String("secret", "s6Ndh+pPbnzHbS*+9Pk8qGwhtzbpa@ge", "Secret Key")
  dsn := flag.String("dsn", "root:e@tcp(db:3306)/seattleGophers?parseTime=true", "MySQL data source name")
  flag.Parse()

  infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
  errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

  db, err := openDB(*dsn)
  if err != nil {
    errorLog.Fatal(err)
  }
  defer db.Close()

  templateCache, err := newTemplateCache("./ui/html/") // Puts template files into cache for easy access
  if err != nil {
    errorLog.Fatal(err)
    }

  session := sessions.New([]byte(*secret))
  session.Lifetime = 12 * time.Hour

  app := &application {
    errorLog:        errorLog,
    infoLog:         infoLog,
    templateCache:   templateCache,
    session:         session,
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
  go http.ListenAndServe(":8080", http.HandlerFunc(redirect)) 
  err = srv.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem")
  errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
  db, err := sql.Open("mysql", dsn)
  if err != nil {
    return nil, err
  }
  if err = db.Ping(); err != nil {
    return nil, err
  }
  return db, nil
}
func redirect(w http.ResponseWriter, r *http.Request) {
  target := "https://" + r.Host + r.URL.Path
  if len(r.URL.RawQuery) > 0 {
    target += "?" + r.URL.RawQuery
  }
  log.Printf("redirect to: %s", target)
  http.Redirect(w, r, target,
    http.StatusTemporaryRedirect)
}
