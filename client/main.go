package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"text/template"
	"time"
)

type ExchangeRateResponse struct {
	Bid float64 `json:"bid,string"`
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != 200 {
		fmt.Println(string(respBody))
		return
	}

	var exchangeRate ExchangeRateResponse
	json.Unmarshal(respBody, &exchangeRate)

	f, err := os.Create("cotacao.txt")
	if err != nil {
		panic(err)
	}

	tmp := template.New("ExchangeRateTemplate")
	tmp, err = tmp.Parse("Dólar: {{.Bid}}")
	if err != nil {
		panic(err)
	}

	tmp.Execute(f, exchangeRate)
}
