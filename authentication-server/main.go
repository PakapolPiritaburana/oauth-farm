package main

import (
	"context"

	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	a := New()
	a.registerRoutes()
	go a.start()

	stop := make(chan os.Signal, 1)

	signal.Notify(stop, syscall.SIGTERM)

	<-stop

	// pkill -15 main
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	err := a.shutdown(ctx)
	if err != nil {
		panic(err)
	}
}
