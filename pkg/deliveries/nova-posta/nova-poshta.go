package nova_posta

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type novaPoshta struct {
	apiKey string
}

func NewNovaPoshta(apiKey string) *novaPoshta {
	return &novaPoshta{
		apiKey: apiKey,
	}
}


func (np *novaPoshta) TrackingDocument(number, phone string) {
	r := novaPoshtaRequest{
		ModelName:    "TrackingDocument",
		CalledMethod: "getStatusDocuments",
	}
	document := trackingDocument{
		number,
		phone,
	}
	methodProperties := trackingDocuments{}
	methodProperties.Documents = append(methodProperties.Documents, document)
	methodProperties.CheckWeightMethod = "3"
	r.MethodProperties = methodProperties

	np.makeRequest(r)
}

func (np *novaPoshta) makeRequest(r novaPoshtaRequest) interface{} {
	r.ApiKey = np.apiKey

	data, err := json.Marshal(r)

	if err != nil {
		log.Fatal(err)
	}
	reader := bytes.NewReader(data)

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, "https://api.novaposhta.ua/v2.0/json/", reader)

	if err != nil {
		fmt.Println(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}


	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bytes))

	npResponse:= &novaPoshtaResponse{}

	err = json.Unmarshal(bytes, npResponse)

	fmt.Println(resp.Status)
	fmt.Println(npResponse)
	fmt.Println(npResponse.Success)
	fmt.Println(npResponse.Data)


	return 5
}
