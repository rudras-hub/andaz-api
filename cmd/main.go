package main

import (
	"andaz-api/internal/database/adapter"
	"andaz-api/internal/logging"
	"andaz-api/internal/usecase"
	"context"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"os/signal"
	"time"
)

const(
	serviceName = "[andaz-api] "
	serviceAddr = ":9090"
	characterGetPath = "/characters"
	readTimeoutSec = 5
	writeTimeoutSec = 10
	IdleTimeoutSec = 120
	shutDownWaitTimeSec = 30
)
func main(){
	logExt := logging.NewStdLoggerExtension(serviceName)
	router := mux.NewRouter()

	charRepo := adapter.GetCharacterRepositoryInstance()
	charGetHandler := usecase.NewCharacterGetHandler(logExt, charRepo)

	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc(characterGetPath, charGetHandler.ListAll)

	s := http.Server{
		Addr:         serviceAddr,       // configure the bind address
		Handler:      router,            // set the default handler
		ErrorLog:     logExt.Logger,     // set the logger for the server
		ReadTimeout:  readTimeoutSec * time.Second,   // max time to read request from the client
		WriteTimeout: writeTimeoutSec * time.Second,  // max time to write response to the client
		IdleTimeout:  IdleTimeoutSec * time.Second, // max time for connections using TCP Keep-Alive
	}

	// start the server
	go func() {
		logExt.Infof("Starting server on port %v", serviceAddr)

		err := s.ListenAndServe()
		if err != nil {
			logExt.Errorf("Error starting server: %v", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	logExt.Infof("Got signal: %v", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), shutDownWaitTimeSec*time.Second)
	s.Shutdown(ctx)
}