package dhl

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFixturesTrackingDocument(t *testing.T) {
	testCases := []struct {
		name        string
		file        string
		trackNumber string
		status      int
		err         error
	}{
		{
			name:        "Tracked by number 1",
			file:        "fixtures/tracked_by_number_1.json",
			trackNumber: "7777777770",
			status:      http.StatusOK,
		},
		{
			name:        "Tracked by number 2",
			file:        "fixtures/tracked_by_number_2.json",
			trackNumber: "00340434292135100186",
			status:      http.StatusOK,
		},
		{
			name:        "Invalid input",
			file:        "fixtures/invalid_input.json",
			trackNumber: "",
			status:      http.StatusBadRequest,
			err:         errors.New("Invalid input: missing mandatory parameter 'trackingNumber'"),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(testCase.status)

				b, err := ioutil.ReadFile(testCase.file)
				if err != nil {
					t.Fatal(err)
				}

				_, err = w.Write(b)
				assert.NoError(t, err)
			}))
			defer server.Close()

			dhl := NewApi(server.URL, "")

			res, err := dhl.TrackingDocument(testCase.trackNumber)
			assert.Equal(t, testCase.err, err)
			if res != nil {
				assert.Equal(t, testCase.trackNumber, res.Shipments[0].Id)
			}
		})
	}
}
