package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	aplistener "github.com/vallinplasencia/boletia/v1/pkg/listener"
	apserver "github.com/vallinplasencia/boletia/v1/pkg/server"
)

func main() {
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// server up
	s := apserver.New()
	s.Run()

	// listener up
	l := aplistener.New()
	l.Run()

	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	s.Shutdown(ctx)
	l.Shutdown(ctx)
}
