package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"time"

	common "mongodb.com/common/application"
)

func server() context.CancelFunc {
	port := os.Getenv("PORT")

	server := &http.Server{
		Addr:         port,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      routes(),
	}

	go func() {
		common.InfoLog.Println("Server running on port", port)
		if err := server.ListenAndServe(); err != nil {
			common.ErrorLog.Println(err)
		}
	}()

	var wait time.Duration

	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")

	flag.Parse()

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), wait)

	defer cancel()

	server.Shutdown(ctx)

	common.InfoLog.Println("Server stopped")

	os.Exit(0)

	return cancel

}
