package np

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrackingDocument(t *testing.T) {
	np := NewNovaPoshta("https://api.novaposhta.ua", "")

	document := TrackingDocument{
		DocumentNumber: "445",
		Phone:          "",
	}
	methodProperties := TrackingDocuments{}
	methodProperties.Documents = append(methodProperties.Documents, document)
	methodProperties.CheckWeightMethod = "3"

	data, err := np.TrackingDocument(methodProperties)
	fmt.Println(data, err)
}

func TestTrackingDocument2(t *testing.T) {
	testCases := []struct {
		name     string
		file     string
		document TrackingDocument
		error    bool
	}{
		{
			name: "Tracked by number",
			file: "fixtures/tracked_by_number.json",
			document: TrackingDocument{
				DocumentNumber: "",
				Phone:          "",
			},
			error: false,
		},
		{
			name: "Tracked by number",
			file: "fixtures/tracked_by_number_and_phone.json",
			document: TrackingDocument{
				DocumentNumber: "",
				Phone:          "",
			},
			error: false,
		},
		{
			name: "Tracked by number",
			file: "fixtures/invalid_number.json",
			document: TrackingDocument{
				DocumentNumber: "",
				Phone:          "",
			},
			error: true,
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

			np := NewNovaPoshta(server.URL, "")

			methodProperties := TrackingDocuments{}
			methodProperties.Documents = append(methodProperties.Documents, testCase.document)
			methodProperties.CheckWeightMethod = "3"

			data, err := np.TrackingDocument(methodProperties)
			assert.NoError(t, err)

			if len(data.Errors) > 0 {
				assert.True(t, testCase.error)
			}
		})
	}
}
