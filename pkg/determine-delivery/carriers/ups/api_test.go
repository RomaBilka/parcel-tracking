package ups

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	response_errors "github.com/RomaBilka/parcel-tracking/pkg/response-errors"
	"github.com/stretchr/testify/assert"
)

func TestApi_TrackByTrackingNumber(t *testing.T) {
	testCases := []struct {
		name           string
		file           string
		trackingNumber string
		errorCode      string
		err            error
	}{
		{
			name:           "Tracked by number",
			file:           "fixtures/tracked_by_number.xml",
			trackingNumber: "1Z12345E0291980793",
		},
		{
			name: "Invalid number",
			file: "fixtures/invalid_number.xml",
			err:  response_errors.InvalidNumber,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				b, err := os.ReadFile(testCase.file)
				if err != nil {
					t.Fatal(err)
				}

				_, err = w.Write(b)
				assert.NoError(t, err)
			}))
			defer server.Close()

			ups := NewApi(server.URL, "", "", "")

			res, err := ups.TrackByTrackingNumber(testCase.trackingNumber)
			assert.Equal(t, testCase.err, err)
			if res != nil {
				assert.Equal(t, testCase.trackingNumber, res.Shipment.ShipmentIdentificationNumber)
			}
		})
	}
}
