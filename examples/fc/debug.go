package main

import (
	"errors"
	"log"
	"net/http"
	"net/http/pprof"
)

func StartEventServer() error {
	mux := http.NewServeMux()

	// Register pprof handlers
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)

	addr := "0.0.0.0:23230"
	server := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Printf("Failed to start event server: %v", err)
		return err
	}
	return nil
}
