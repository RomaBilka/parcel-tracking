package usps

import (
	"encoding/xml"
	"errors"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsError(t *testing.T) {
	resp := &response{}
	testCases := []struct {
		name string
		file string
		err  error
	}{
		{
			name: "No error",
			file: "fixtures/track_response.xml",
			err:  nil,
		},
		{
			name: "bad xml, parsing error",
			err: &xml.SyntaxError{
				Msg:  "element <Bad> closed by </BadSyntax>",
				Line: 1,
			},
		},
		{
			name: "xml body contain auth error",
			file: "fixtures/error_response_do_auth.xml",
			err:  errors.New("USPS body error.\n Number: 80040B1A\n Description: API Authorization failure. User XXXXXXX is not authorized to use API TrackV2.\n Source: USPSCOM::DoAuth\n HelpContext: \n"),
		},
		{
			name: "xml body contain missing value error",
			file: "fixtures/error_response_missing_value.xml",
			err:  errors.New("USPS body error.\n Number: -2147217951\n Description: Missing value for To Phone number.\n Source: EMI_Respond :EMI:clsEMI.ValidateParameters:\n        clsEMI.ProcessRequest;SOLServerIntl.EMI_Respond\n    \n HelpContext: 1000440\n"),
		},
	}

	var (
		xmlBody []byte
		err     error
	)

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.file == "" {
				xmlBody = []byte(`<Bad></BadSyntax>`)
			} else {
				xmlBody, err = ioutil.ReadFile(testCase.file)
				if err != nil {
					t.Error(err)
				}
			}

			err = resp.isError(xmlBody)
			assert.Equal(t, testCase.err, err)
		})
	}
}

func TestUnmarshalTrackData(t *testing.T) {

	testCases := []struct {
		name    string
		file    string
		number  string
		details []string
		err     error
	}{
		{
			name:   "No error",
			file:   "fixtures/track_response.xml",
			number: "EJ958083578US",
			details: []string{
				"Your item was delivered at 8:10 am on June 1 in Wilmington DE 19801.",
				"May 30 11:07 am NOTICE LEFT WILMINGTON DE 19801.",
				"May 30 10:08 am ARRIVAL AT UNIT WILMINGTON DE 19850.",
				"May 29 9:55 am ACCEPT OR PICKUP EDGEWATER NJ 07020.",
			},
			err: nil,
		},
		{
			name: "bad xml, parsing error",
			err: &xml.SyntaxError{
				Msg:  "element <Bad> closed by </BadSyntax>",
				Line: 1,
			},
		},
		{
			name: "xml body contain auth error",
			file: "fixtures/error_response_do_auth.xml",
			err:  errors.New("the XML structure does not match the parsing conditions or is missing"),
		},
	}

	var (
		xmlBody []byte
		err     error
	)

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.file == "" {
				xmlBody = []byte(`<Bad></BadSyntax>`)
			} else {
				xmlBody, err = ioutil.ReadFile(testCase.file)
				if err != nil {
					t.Error(err)
				}
			}

			resp := response{}

			err = resp.unmarshalTrackData(xmlBody)
			assert.Equal(t, testCase.err, err)
			assert.Equal(t, testCase.number, resp.number)
		})
	}
}
