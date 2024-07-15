package imposter_bank

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/cko-recruitment/payment-gateway-challenge-go/internal/logger"
	"github.com/cko-recruitment/payment-gateway-challenge-go/internal/models/dtos"
	"io"
	"net/http"
)

// HTTPClient helps for mocking httpclien
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// ConnectorConfig used to store all configs related to bank.
type ConnectorConfig struct {
	HTTPClient HTTPClient
	BankURL    string
}

type Connector struct {
	httpClient HTTPClient
	BankURL    string
}

func NewImposterBankConnector(config ConnectorConfig) *Connector {
	return &Connector{
		httpClient: config.HTTPClient,
		BankURL:    config.BankURL,
	}
}

func (c *Connector) MakePayment(ctx context.Context, requestBody dtos.BankPaymentRequest) (*dtos.BankPaymentResponse, error) {
	jsonDataBankReq, _ := json.Marshal(requestBody)
	var bankResponse dtos.BankPaymentResponse
	_ = json.Unmarshal(jsonDataBankReq, &bankResponse)

	req, err := http.NewRequestWithContext(ctx, "POST", c.BankURL, bytes.NewBuffer(jsonDataBankReq))

	if err != nil {
		return nil, err
	}
	req = req.WithContext(req.Context())
	rsp, err := c.httpClient.Do(req)

	defer func() {
		if err := rsp.Body.Close(); err != nil {
			logger.Error(fmt.Sprintf("cannot close request %s", err))
		}
	}()

	if err != nil {
		logger.Error(fmt.Sprintf("error in reading response body %s", err))
		return nil, fmt.Errorf("error in reading response body: %w", err)
	}

	byteResponse, err := io.ReadAll(rsp.Body)

	// check response
	if rsp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bank status code error with %d", rsp.StatusCode)
	}

	err = json.Unmarshal(byteResponse, &bankResponse)

	return &bankResponse, nil
}
