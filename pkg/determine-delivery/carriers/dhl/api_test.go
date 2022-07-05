package dhl

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrackingDocument(t *testing.T) {
	dhl := NewApi("https://api-eu.dhl.com", "demo-key")

	fmt.Println(dhl.TrackingDocument("00340434292135100186"))
}

func TestFixturesTrackingDocument(t *testing.T) {
	testCases := []struct {
		name        string
		file        string
		trackNumber string
	}{
		{
			name:        "Tracked by number 1",
			file:        "fixtures/tracked_by_number_1.json",
			trackNumber: "",
		},
		{
			name:        "Tracked by number 2",
			file:        "fixtures/tracked_by_number_2.json",
			trackNumber: "",
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)

				b, err := ioutil.ReadFile(testCase.file)
				if err != nil {
					t.Fatal(err)
				}

				_, err = w.Write(b)
				assert.NoError(t, err)
			}))
			defer server.Close()

			fmt.Println(server.URL)
			dhl := NewApi(server.URL, "")

			_, err := dhl.TrackingDocument(testCase.trackNumber)
			assert.NoError(t, err)
		})
	}
}
