package quickbooks

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/syngency/go-quickbooks.v2/sdk"
	"github.com/syngency/go-quickbooks.v2/sdk/consts"
)

// Quickbooks client type
type Quickbooks struct {
	RealmID     string
	AccessToken string
	baseURL     string
}

// Address quickbooks address object
type Address struct {
	ID                     string `json:"Id,omitempty"`
	Line1                  string `json:"Line1"`
	Line2                  string `json:"Line2,omitempty"`
	City                   string `json:"City"`
	CountrySubDivisionCode string `json:"CountrySubDivisionCode"`
	PostalCode             string `json:"PostalCode"`
	Lat                    string `json:"Lat,omitempty"`
	Long                   string `json:"Long,omitempty"`
}

// NewClient creates a new client to work with Quickbooks
func NewClient(realmID string, accessToken string, isSandbox bool) *Quickbooks {
	q := Quickbooks{}
	q.RealmID = realmID
	q.AccessToken = accessToken

	if isSandbox {
		q.baseURL = sdk.SandboxURL
	} else {
		q.baseURL = sdk.ProductionURL
	}

	return &q
}

// makeGetRequest makes a GET request to Quickbooks API.
// endpoint should start with a leading '/'
func (q *Quickbooks) makeGetRequest(endpoint string) (*http.Response, error) {
	urlStr := q.baseURL + endpoint
	httpClient := &http.Client{}

	request, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		return nil, err
	}

	// headers
	request.Header.Set("accept", "application/json")
	request.Header.Set("Authorization", "Bearer "+q.AccessToken)

	response, err := httpClient.Do(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, handleError(response)
	}

	return response, nil
}

// PostRequest makes a POST request to Quickbooks API.
// endpoint should start with a leading '/'
func (q *Quickbooks) makePostRequest(endpoint string, body interface{}) (*http.Response, error) {
	urlStr := q.baseURL + endpoint
	httpClient := &http.Client{}

	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", urlStr, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}

	// headers
	request.Header.Set("accept", "application/json")
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+q.AccessToken)

	response, err := httpClient.Do(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, handleError(response)
	}

	return response, nil
}

func handleError(response *http.Response) error {
	switch response.StatusCode {
	case 400:
		qbError := ErrorObject{}
		err := json.NewDecoder(response.Body).Decode(&qbError)
		if err != nil {
			return err
		}

		return qbError
	case 401:
		sdkError := SDKError{}
		return sdkError.New(consts.QBAuthorizationFault, consts.QBAuthenticationFaultCode, consts.QBAuthorizationFaultMessage)
	}

	return nil
}
