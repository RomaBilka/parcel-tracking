package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/RomaBilka/parcel-tracking/api/rest"
	"github.com/RomaBilka/parcel-tracking/dependencies"
	"golang.org/x/sync/errgroup"
)

func main() {
	deps, err := dependencies.InitDeps()
	if err != nil {
		panic(err)
	}
	defer deps.TearDown()

	rest.Configure(deps)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	httpServer := &http.Server{
		Addr: ":" + deps.Config.Port,
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
