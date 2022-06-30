package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/RomaBilka/parcel-tracking/cmd"
	"github.com/RomaBilka/parcel-tracking/internal/handlers"
	"golang.org/x/sync/errgroup"
)

func main() {
	tracker := handlers.NewTracker(cmd.GetDetector())
	http.HandleFunc("/tracking", tracker.Tracking)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	httpServer := &http.Server{
		Addr: ":" + cmd.O.Port,
	}
	g, gCtx := errgroup.WithContext(ctx)
	g.Go(func() error {
		fmt.Println("Server is listening...")
		return httpServer.ListenAndServe()
	})
	g.Go(func() error {
		<-gCtx.Done()
		return httpServer.Shutdown(context.Background())
	})
	if err := g.Wait(); err != nil {
		fmt.Printf("exit reason: %s \n", err)
	}
}
