package service

import (
	"context"
	"encoding/json"
	"finance/api/models"
	"finance/storage"
	"fmt"
	"net/http"
)

type currencyService struct {
	storage storage.IStorage
}


func NewCurrencyService(storage storage.IStorage) currencyService {
	return currencyService{storage: storage}
}

type ExchangeRateAPIResponse struct {
	Data map[string]struct {
		Value float64 `json:"value"`
	} `json:"data"`
}

func (c currencyService) GetRate(ctx context.Context,apiKey string,req models.ExchangeRateRequest) (string, error) {
    url :=fmt.Sprintf("https://api.currencyapi.com/v3/latest?apikey=%s&currencies=%s&base_currency=%s", apiKey,req.To,req.From)

    resp, err := http.Get(url)
    if err != nil {
        return "0", fmt.Errorf("error making HTTP request: %w", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return "0", fmt.Errorf("received response code: %d", resp.StatusCode)
    }

    var response ExchangeRateAPIResponse
    if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
        return "0", fmt.Errorf("error decoding JSON response: %w", err)
    }

    currencyRate, ok := response.Data[req.To]
    if !ok {
        return "0", fmt.Errorf("rate not found in response")
    }

    if req.Amount==0 {
        req.Amount=1
    }

    msg:=fmt.Sprintf("%d %s==%d %s",req.Amount,req.From,int(currencyRate.Value)*req.Amount,req.To)

    return msg, nil
}