package np

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
		name     string
		file     string
		document TrackingDocument
		error    error
	}{
		{
			name: "Tracked by number",
			file: "fixtures/tracked_by_number.json",
			document: TrackingDocument{
				DocumentNumber: "59000777777777",
				Phone:          "",
			},
			error: nil,
		},
		{
			name: "Tracked by number and phone",
			file: "fixtures/tracked_by_number_and_phone.json",
			document: TrackingDocument{
				DocumentNumber: "59000777777777",
				Phone:          "",
			},
			error: nil,
		},
		{
			name: "Invalid number",
			file: "fixtures/invalid_number.json",
			document: TrackingDocument{
				DocumentNumber: "",
				Phone:          "",
			},
			error: errors.New("Document number is not correct"),
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

			methodProperties := TrackingDocuments{}
			methodProperties.Documents = append(methodProperties.Documents, testCase.document)
			methodProperties.CheckWeightMethod = "3"

			res, err := np.TrackingDocument(methodProperties)
			assert.Equal(t, testCase.error, err)
			if res != nil {
				assert.Equal(t, testCase.document.DocumentNumber, res.Data[0].Number)
			}

		})
	}
}
