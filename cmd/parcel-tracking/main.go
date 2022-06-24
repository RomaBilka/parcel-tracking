package main

import (
	"fmt"
	"net/http"

	"github.com/RomaBilka/parcel-tracking/internal/handlers"
	determine_delivery "github.com/RomaBilka/parcel-tracking/pkg/determine-delivery"
	"github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers/me"
	"github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers/np"
)

func main() {
	detector := determine_delivery.NewDetector()
	detector.Registry(np.NewCarrier(np.NewApi("https://api.novaposhta.ua", "")))
	detector.Registry(me.NewCarrier(me.NewApi("0xA79E003048D2B47311E26B7D4A430FFC", "public", "PUBLIC", "https://apii.meest-group.com/T/1C_Query.php")))

	tracker := handlers.NewTracker(detector)
	http.HandleFunc("/tracking", tracker.Tracking)

	fmt.Println("Server is listening...")
	_ = http.ListenAndServe(":8080", nil)
}
