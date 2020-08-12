package app

import (
	"cyberuranus-interview/pkg/data"
	"database/sql"
	"log"
	"net/http"
)

type httpMode int

// TODO: Add multiple modes for every situation
const (
	debugMode httpMode = iota
	productionMode
)

func Start() {
	db := NewDB()
	initAPIServer(db)
}

func initDebugLogging() {
	// TODO: Control Logging System
}

func initAPIServer(db *sql.DB) {
	mux := http.NewServeMux()
	data.Register(mux, db)
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", mux))
}
