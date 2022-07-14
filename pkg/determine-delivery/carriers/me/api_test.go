package me

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApi_ShipmentsTrack(t *testing.T) {
	api := NewApi("1", "1", "1", "https://apii.meest-group.com/T/1C_Query.php")
	fmt.Println(api.ShipmentsTrack("dddd"))
}

func TestFixturesTrackingDocument(t *testing.T) {
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
			name:      "Tracked by number",
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

				b, err := ioutil.ReadFile(testCase.file)
				if err != nil {
					t.Fatal(err)
				}

				_, err = w.Write(b)
				assert.NoError(t, err)
			}))
			defer server.Close()

			me := NewApi("0xA79E003048D2B47311E26B7D4A430FFC", "public", "PUBLIC", server.URL)

			r, err := me.ShipmentsTrack(testCase.document)
			if err == nil {
				assert.Equal(t, testCase.errorCode, r.Errors.Code)
			}
			assert.Equal(t, testCase.err, err)
		})
	}
}
