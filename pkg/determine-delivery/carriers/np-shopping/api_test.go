package np_shopping

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFixturesTrackingDocument(t *testing.T) {
	testCases := []struct {
		name          string
		file          string
		documentID    string
		errorContains string
	}{
		{
			name:       "Tracked by number",
			file:       "fixtures/tracked_by_number.json",
			documentID: "EV00000000000543NPI",
		},
		{
			name:          "Invalid number",
			file:          "fixtures/invalid_number.json",
			errorContains: "document number is not correct",
		},
		{
			name:          "Internal error",
			file:          "fixtures/invalid_response.json",
			errorContains: `something went wrong: {"invalid_field" : "invalid_value"}`,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				b, err := os.ReadFile(tc.file)
				if err != nil {
					t.Fatal(err)
				}

				_, err = w.Write(b)
				assert.NoError(t, err)
			}))
			defer server.Close()

			np := Api{url: server.URL}
			res, err := np.TrackingDocument(tc.documentID)
			if tc.errorContains != "" {
				assert.ErrorContains(t, err, tc.errorContains)
			} else {
				assert.NoError(t, err)
			}

			if tc.documentID == "" {
				assert.Nil(t, res)
			} else {
				assert.Equal(t, tc.documentID, res.WaybillNumber)
			}
		})
	}
}
