package client

import (
	"encoding/json"
	"fmt"
	"github.com/yigitsadic/qrmenum_client/shared"
	"net/http"
)

type HTTPClient struct {
	BaseUrl string
}

func (h HTTPClient) FetchFromCMS(key string) ([]ProductResponse, error) {
	resp, err := http.Get(fmt.Sprintf("%s/products?company.label=%s", h.BaseUrl, key))

	if err != nil {
		return nil, shared.UnableToFetchFromCMS
	}

	var values []ProductResponse

	err = json.NewDecoder(resp.Body).Decode(&values)
	if err != nil {
		return nil, shared.UnableMarshalResponseFromCMS
	}

	if len(values) == 0 {
		return nil, shared.KeyNotFoundOnCMS
	}

	return values, nil
}

func NewHTTPClient(baseUrl string) *HTTPClient {
	return &HTTPClient{BaseUrl: baseUrl}
}
