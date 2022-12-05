package fedex

import (
	"errors"
	"testing"

	"github.com/RomaBilka/parcel-tracking/pkg/determine-delivery/carriers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCarrier_Detect(t *testing.T) {
	testCases := []struct {
		name    string
		trackId string
		ok      bool
	}{
		{name: "12 true", trackId: "123456789012", ok: true},
		{name: "15 true", trackId: "123456789012345", ok: true},
		{name: "20 true", trackId: "12345678901234567890", ok: true},
		{name: "22 true", trackId: "1234567890123456789012", ok: true},
		{name: "21 false", trackId: "123456789012345678901", ok: false},
		{name: "10 false", trackId: "1234567890", ok: false},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			fedex := NewCarrier(NewApi("", "", "", ""))
			ok := fedex.Detect(testCase.trackId)
			assert.Equal(t, testCase.ok, ok)
		})
	}
}

func TestCarrier_Track(t *testing.T) {
	testCases := []struct {
		name         string
		trackNumbers []string
		setupApiMock func(api *apiMock, trackNumbers []string)
		parcels      []carriers.Parcel
		err          error
	}{
		{
			name:         "Ok response",
			trackNumbers: []string{"12A12345", "12A12346"},
			setupApiMock: func(api *apiMock, trackNumbers []string) {

				trackingInfo := TrackingInfo{
					TrackingNumberInfo: TrackingNumberInfo{
						TrackingNumber: trackNumbers[0],
					},
				}

				requestData1 := TrackingRequest{IncludeDetailedScans: true}
				requestData1.TrackingInfo = append(requestData1.TrackingInfo, trackingInfo)

				trackingInfo = TrackingInfo{
					TrackingNumberInfo: TrackingNumberInfo{
						TrackingNumber: trackNumbers[1],
					},
				}
				requestData2 := TrackingRequest{IncludeDetailedScans: true}
				requestData2.TrackingInfo = append(requestData2.TrackingInfo, trackingInfo)

				response1 := &TrackingResponse{
					Output: Output{
						CompleteTrackResults: []CompleteTrackResult{
							CompleteTrackResult{
								TrackingNumber: trackNumbers[0],
								TrackResults: []TrackResult{
									TrackResult{
										LatestStatusDetail: LatestStatusDetail{
											StatusByLocale: "Ok",
										},
										OriginLocation: OriginLocation{
											LocationContactAndAddress: ContactAndAddress{Address: Address{}},
										},
										DeliveryDetails: DeliveryDetails{ActualDeliveryAddress: Address{}},
										RecipientInformation: ContactAndAddress{
											Address: Address{
												CountryName: "Country Name",
												City:        "City",
											},
										},
									},
								},
							},
						},
					},
				}

				response2 := &TrackingResponse{
					Output: Output{
						CompleteTrackResults: []CompleteTrackResult{
							CompleteTrackResult{
								TrackingNumber: trackNumbers[1],
								TrackResults: []TrackResult{
									TrackResult{
										LatestStatusDetail: LatestStatusDetail{
											StatusByLocale: "Ok",
										},
										OriginLocation: OriginLocation{
											LocationContactAndAddress: ContactAndAddress{Address: Address{}},
										},
										DeliveryDetails: DeliveryDetails{ActualDeliveryAddress: Address{}},
										RecipientInformation: ContactAndAddress{
											Address: Address{
												CountryName: "Country Name",
												City:        "City",
											},
										},
									},
								},
							},
						},
					},
				}

				api.
					On("TrackByTrackingNumber", requestData1).
					Once().Return(response1, nil).
					On("TrackByTrackingNumber", requestData2).
					Once().Return(response2, nil)
			},
			parcels: []carriers.Parcel{
				carriers.Parcel{TrackingNumber: "12A12345", Places: []carriers.Place{carriers.Place{Country: "Country Name", City: "City"}}, Status: "Ok"},
				carriers.Parcel{TrackingNumber: "12A12346", Places: []carriers.Place{carriers.Place{Country: "Country Name", City: "City"}}, Status: "Ok"},
			},
		},
		{
			name:         "Bad response",
			trackNumbers: []string{""},
			setupApiMock: func(api *apiMock, trackNumbers []string) {
				trackingInfo := TrackingInfo{
					TrackingNumberInfo: TrackingNumberInfo{
						TrackingNumber: trackNumbers[0],
					},
				}

				requestData := TrackingRequest{IncludeDetailedScans: true}
				requestData.TrackingInfo = append(requestData.TrackingInfo, trackingInfo)
				api.On("TrackByTrackingNumber", requestData).Once().Return(nil, errors.New("bad request"))
			},
			err: errors.New("bad request"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			api := &apiMock{}
			testCase.setupApiMock(api, testCase.trackNumbers)

			c := NewCarrier(api)
			parcels, err := c.Track(testCase.trackNumbers)

			assert.Equal(t, testCase.err, err)
			assert.ElementsMatch(t, testCase.parcels, parcels)

			api.AssertExpectations(t)
		})
	}
}

type apiMock struct {
	mock.Mock
}

func (m *apiMock) TrackByTrackingNumber(trackingRequest TrackingRequest) (*TrackingResponse, error) {
	arg := m.Called(trackingRequest)
	if arg.Get(0) == nil {
		return nil, arg.Error(1)
	}

	return arg.Get(0).(*TrackingResponse), arg.Error(1)
}
