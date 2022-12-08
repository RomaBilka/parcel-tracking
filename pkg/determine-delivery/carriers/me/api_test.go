package me

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApi_TrackByTrackingNumber(t *testing.T) {
	testCases := []struct {
		name      string
		file      string
		document  string
		errorCode string
		err       error
	}{
		{
			name:      "Tracked by number",
			file:      "fixtures/tracked_by_number.xml",
			document:  "TESTIK11",
			errorCode: "000",
		},
		{
			name:      "Bad request",
			file:      "fixtures/bad_request.xml",
			document:  "TESTIK11",
			errorCode: "101",
			err:       errors.New("Connection Error"),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)

				b, err := os.ReadFile(testCase.file)
				if err != nil {
					t.Fatal(err)
				}

				_, err = w.Write(b)
				assert.NoError(t, err)
			}))
			defer server.Close()

			me := NewApi(server.URL, "0xA79E003048D2B47311E26B7D4A430FFC", "public", "PUBLIC")

			res, err := me.TrackByTrackingNumber(testCase.document)
			assert.Equal(t, testCase.err, err)
			if res != nil {
				assert.Equal(t, testCase.errorCode, res.Errors.Code)
			}
		})
	}
}
