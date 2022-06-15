package me

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrackingDocument(t *testing.T) {
	me := NewApi("0xA79E003048D2B47311E26B7D4A430FFC", "public", "PUBLIC", "https://apii.meest-group.com/T/1C_Query.php")

	_, err := me.ShipmentsTrack("TESTIK11")
	assert.NoError(t, err)
}

func TestFixturesTrackingDocument(t *testing.T) {
	testCases := []struct {
		name      string
		file      string
		document  string
		errorCode string
	}{
		{
			name:      "Tracked by number",
			file:      "fixtures/tracked_by_number.xml",
			document:  "TESTIK11",
			errorCode: "000",
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

				w.Write(b)
			}))
			defer server.Close()

			me := NewApi("0xA79E003048D2B47311E26B7D4A430FFC", "public", "PUBLIC", server.URL)

			r, err := me.ShipmentsTrack(testCase.document)
			assert.NoError(t, err)
			assert.Equal(t, testCase.errorCode, r.Errors.Code)
		})
	}
}
