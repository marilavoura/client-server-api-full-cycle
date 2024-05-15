package services

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"server/models"
	"time"
)

type ExchangeRateResponse struct {
	UsdBrl models.ExchangeRate `json:"USDBRL"`
}

func GetExchangeRate(parentContext context.Context) (*ExchangeRateResponse, error) {
	ctx, cancel := context.WithTimeout(parentContext, 200*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	exchangeRateJson, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var exchangeRate ExchangeRateResponse
	json.Unmarshal(exchangeRateJson, &exchangeRate)

	return &exchangeRate, nil
}
