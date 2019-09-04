package main

import (
  "flag"
  "log"
  "crypto/tls"
  "os"
  "net/http"
  )




func main() {
//standard development port = 3333 - at runtime use flag to set to https port 443
  addr := flag.String("addr", ":3333", "Port to accept incoming connections")
  flag.Parse()

//standard error logging.  Can be piped at runtime to write to a file.
  infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
  errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

//set server to use more secure protocols
  tlsConfig := &tls.Config {
    PreferServerCipherSuites: true,
    CurvePreferences:         []tls.CurveID{tls.X25519, tls.CurveP256},
    }
  srv := &http.Server {
    Addr:         *addr,
    ErrorLog:     errorLog,
    Handler:      routes(), // see cmd/web/routes.go
    TLSConfig:    tlsConfig,
    }
//start TLS server
  infoLog.Printf("Starting server on %s", *addr)
  //for testing generate a self-signed cert with https://golang.org/src/crypto/tls/generate_cert.go
  //   $ cd tls; go run /usr/local/go/src/crypto/tls/generate_cert.go --rsa-bits=2048 --host=localhost
  err := srv.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem")
  errorLog.Fatal(err)
  }
