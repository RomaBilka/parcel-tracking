package ups

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApi_TrackByNumber(t *testing.T) {
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
			err:  errors.New("Invalid tracking number"),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				b, err := ioutil.ReadFile(testCase.file)
				if err != nil {
					t.Fatal(err)
				}

				_, err = w.Write(b)
				assert.NoError(t, err)
			}))
			defer server.Close()

			ups := NewApi(server.URL, "", "", "")

			res, err := ups.TrackByNumber(testCase.trackingNumber)
			assert.Equal(t, testCase.err, err)
			if res != nil {
				assert.Equal(t, testCase.trackingNumber, res.Shipment.ShipmentIdentificationNumber)
			}
		})
	}
}
