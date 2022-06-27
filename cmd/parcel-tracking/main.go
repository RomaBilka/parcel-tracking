package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"

	"github.com/RomaBilka/parcel-tracking/internal/handlers"
	determine_delivery "github.com/RomaBilka/parcel-tracking/pkg/determine-delivery"
	"github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers/me"
	"github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers/np"
	"github.com/jessevdk/go-flags"
)

var opts struct {
	Port       string `short:"p" long:"port" description:"Port" required:"true" default:"8080" env:"PORT"`
	NP_API_URL string `long:"NP_API_URL" description:"nova poshta API URL" required:"true" default:"https://api.novaposhta.ua" env:"NP_API_URL"`
	NP_API_Key string `long:"NP_API_Key" description:"nova poshta API key"  default:"" env:"NP_API_KEY"`

	ME_API_URL  string `long:"ME_API_URL" description:"meest express API URL" required:"true" default:"https://apii.meest-group.com/T/1C_Query.php" env:"ME_API_URL"`
	ME_Login    string `long:"ME_Login" description:"meest express login" required:"true" default:"public" env:"ME_LOGIN"`
	ME_Password string `long:"ME_Password" description:"meest express password" required:"true" default:"PUBLIC" env:"ME_PASSWORD"`
	ME_ID       string `long:"ME_ID" description:"meest express ID" required:"true" default:"0xA79E003048D2B47311E26B7D4A430FFC" env:"ME_ID"`
}

func main() {
	o := opts
	_, err := flags.Parse(&o)
	if err != nil {
		panic(err)
	}

	detector := determine_delivery.NewDetector()
	detector.Registry(np.NewCarrier(np.NewApi(o.NP_API_URL, o.NP_API_Key)))
	detector.Registry(me.NewCarrier(me.NewApi(o.ME_ID, o.ME_Login, o.ME_Password, o.ME_API_URL)))

	tracker := handlers.NewTracker(detector)
	http.HandleFunc("/tracking", tracker.Tracking)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	httpServer := &http.Server{
		Addr: ":" + o.Port,
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
