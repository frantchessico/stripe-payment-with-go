package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/paymentintent"
)

const stripeSecretKey = "your_secret_key_here"

type PaymentRequest struct {
    Amount int64 `json:"amount"`
    
}

func main() {
    stripe.Key = stripeSecretKey

    http.HandleFunc("/create-payment-intent", handleCreatePaymentIntent)

    fmt.Println("Server started on port 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleCreatePaymentIntent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    // Ler o corpo da requisição
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Decodificar o corpo da requisição em uma estrutura de pagamento
    var paymentRequest PaymentRequest
    err = json.Unmarshal(body, &paymentRequest)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
	

    params := &stripe.PaymentIntentParams{
        Amount:   stripe.Int64(paymentRequest.Amount), // Valor em centavos
        Currency: stripe.String("usd"), // Moeda
    }

    intent, err := paymentintent.New(params)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Retorne a resposta em JSON
    response := map[string]interface{}{
        "message":          "Payment intention created successfully",
        "payment_intent_id": intent.ID,
		"amount": intent.Amount,
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}
