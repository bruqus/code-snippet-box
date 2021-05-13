package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type Config struct {
	Addr string
	StaticDir string
}

type application struct {
	errorLog *log.Logger
	infoLog *log.Logger
}

func main() {
	cfg := new(Config)
	flag.StringVar(&cfg.Addr, "addr", ":4000", "HTTP network address")
	flag.StringVar(&cfg.StaticDir, "static", "./ui/static/", "Path to static assets")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Llongfile)

	app := &application{
		infoLog: infoLog,
		errorLog: errorLog,
	}

	srv := &http.Server{
		Addr: cfg.Addr,
		ErrorLog: errorLog,
		Handler: app.routes(),
	}

	infoLog.Printf("Starting Server on %s", cfg.Addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
