package nova_poshta

import (
	"fmt"
	"testing"
)

func TestTrackingDocument(t *testing.T) {
	np := NewNovaPoshta("")

	document := TrackingDocument{
		DocumentNumber: "",
	}
	methodProperties := TrackingDocuments{}
	methodProperties.Documents = append(methodProperties.Documents, document)
	methodProperties.CheckWeightMethod = "3"

	data, err := np.TrackingDocument(methodProperties)
	fmt.Println(data, err)

}
