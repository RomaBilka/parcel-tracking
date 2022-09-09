package usps

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

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
			trackingNumber: "9400100000000000000000",
		},
		{
			//TODO: check error
			name:           "Error",
			file:           "fixtures/error_response.xml",
			trackingNumber: "9400100000000000000000",
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

			api := NewApi(server.URL, "", "")
			res, err := api.TrackByTrackingNumber([]TrackID{})
			assert.Equal(t, testCase.err, err)
			if res != nil {
				assert.Equal(t, testCase.trackingNumber, res.TrackInfo[0].TrackId)
			}
		})
	}
}
