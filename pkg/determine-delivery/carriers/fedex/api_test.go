package fedex

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	response_errors "github.com/RomaBilka/parcel-tracking/pkg/response-errors"
	"github.com/stretchr/testify/assert"
)

func TestApi_TrackByTrackingNumber(t *testing.T) {
	fedex := NewApi("", "", "", "")
	fedex.token.token = "test_token"
	fedex.token.expire = time.Now().Local().Add(60 * time.Minute)

	testCases := []struct {
		name        string
		file        string
		trackNumber string
		status      int
		err         error
	}{
		{
			name:        "Tracked bu number",
			file:        "fixtures/tracked_bu_number.json",
			status:      http.StatusOK,
			trackNumber: "123456789012",
		},
		{
			name:   "Bad requesr",
			file:   "fixtures/bad_request_error.json",
			status: http.StatusOK,
			err:    errors.New("The given JWT is invalid. Please modify your request and try again."),
		},
		{
			name:   "Tracking tracking number notfound",
			file:   "fixtures/tracking_trackingnumber_notfound.json",
			status: http.StatusOK,
			err:    response_errors.NotFound,
		},
		{
			name:   "Validation errcd wspod",
			file:   "fixtures/validation_errcd_wspod.json",
			status: http.StatusOK,
			err:    errors.New("We are unable to process this shipment for the moment. Try again later or contact FedEx Customer Service."),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(testCase.status)

				b, err := os.ReadFile(testCase.file)
				if err != nil {
					t.Fatal(err)
				}

				_, err = w.Write(b)
				assert.NoError(t, err)
			}))
			defer server.Close()

			fedex.apiURL = server.URL

			res, err := fedex.TrackByTrackingNumber(TrackingRequest{})
			assert.Equal(t, testCase.err, err)
			if res != nil {
				assert.Equal(t, testCase.trackNumber, res.Output.CompleteTrackResults[0].TrackingNumber)
			}
		})
	}
}

func TestApi_authorize(t *testing.T) {
	fedex := NewApi("", "", "", "")

	testCases := []struct {
		name  string
		file  string
		token string
		err   error
	}{
		{
			name:  "Not authorized",
			file:  "fixtures/not_authorized_error.json",
			token: "",
			err:   errors.New("Access token expired. Please modify your request and try again."),
		},
		{
			name:  "Short_lifetime",
			file:  "fixtures/token_with_short_lifetime.json",
			token: "",
			err:   errors.New("short expiration of the token"),
		},
		{
			name:  "Ok",
			token: "test token",
			file:  "fixtures/authorize_ok.json",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				b, err := os.ReadFile(testCase.file)
				if err != nil {
					t.Fatal(err)
				}

				_, err = w.Write(b)
				assert.NoError(t, err)
			}))
			defer server.Close()

			fedex.apiURL = server.URL

			token, err := fedex.authorize()
			assert.Equal(t, testCase.err, err)
			assert.Equal(t, testCase.token, token)
		})
	}
}

func TestApi_isExpired(t *testing.T) {
	testCases := []struct {
		name   string
		expire time.Time
		ok     bool
	}{
		{
			name:   "Ok",
			expire: time.Now().Local().Add(2 * time.Minute),
			ok:     false,
		},
		{
			name:   "expired",
			expire: time.Now().Local().Add(500 * time.Millisecond),
			ok:     true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.ok, isExpired(testCase.expire))
		})
	}
}
