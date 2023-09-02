package main

import (
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	port := ":8000"

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	srv := &http.Server{
		Addr:     port,
		Handler:  app.routes(),
		ErrorLog: errorLog,
	}

	infoLog.Printf("Starting server on port %s ...", port)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
