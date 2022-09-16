package ukrposhta

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApi_TrackByTrackingNumber(t *testing.T) {
	testCases := []struct {
		name     string
		file     string
		barcodes []string
		error    error
	}{
		{
			name: "Tracked by number",
			file: "fixtures/tracked_by_number.json",
			barcodes: []string{
				"0500128254873",
				"0500128254610",
			},
			error: nil,
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

			np := NewApi(server.URL, "")

			res, err := np.TrackByTrackingNumber(testCase.barcodes)
			assert.Equal(t, testCase.error, err)
			for _, b := range testCase.barcodes {
				assert.NotEmpty(t, res.Found[b])
			}
		})
	}
}
