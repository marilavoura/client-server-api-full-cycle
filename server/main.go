package main

import (
	"context"
	"encoding/json"
	"net/http"
	"server/database"
	"server/services"
	"time"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/cotacao", ExchangeRateHandler)
	http.ListenAndServe(":8080", mux)
}

type ExchangeRateHandlerResponse struct {
	Bid float64 `json:"bid,string"`
}

func ExchangeRateHandler(w http.ResponseWriter, r *http.Request) {
	exchangeRateResponse, err := services.GetExchangeRate(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	response := ExchangeRateHandlerResponse{
		Bid: exchangeRateResponse.UsdBrl.Bid,
	}

	responseJson, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	db, err := database.Connect()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	result := db.WithContext(ctx).Create(&exchangeRateResponse.UsdBrl)
	err = result.Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	// fmt.Println(result.RowsAffected)

	// // var exchangeRates []models.ExchangeRate
	// // db.Find(&exchangeRates)
	// // for _, rate := range exchangeRates {
	// // 	fmt.Println(rate)
	// // }

	w.Write(responseJson)
}
