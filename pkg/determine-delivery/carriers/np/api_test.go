package np

import (
	"fmt"
	"testing"
)

func TestTrackingDocument(t *testing.T) {
	np := NewNovaPoshta("https://api.novaposhta.ua", "")

	document := TrackingDocument{
		DocumentNumber: "",
		Phone:          "",
	}
	methodProperties := TrackingDocuments{}
	methodProperties.Documents = append(methodProperties.Documents, document)
	methodProperties.CheckWeightMethod = "3"

	data, err := np.TrackingDocument(methodProperties)
	fmt.Println(data, err)

}
