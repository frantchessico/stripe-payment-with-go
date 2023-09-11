
# Stripe Payment API in Go

This is an example of a Go API that allows you to create a payment intent with Stripe. The API is configured to accept POST requests and responds with details about the created payment intent.

## Prerequisites

- [Go](https://golang.org/doc/install) installed in your environment.
- Stripe secret key. Replace `"your_secret_key_here"` with your Stripe secret key.

## Configuration

1. Install the Stripe Go v72 library:

   ```shell
   go get github.com/stripe/stripe-go/v72
   ```

2. Run the server:

   ```shell
   go run main.go
   ```

The server will be running on port 8080.

## Endpoints

### 1. Create Payment Intent

- **URL:** `/create-payment-intent`
- **Method:** `POST`
- **JSON Request Body:** 

  ```json
  {
      "amount": 1000
  }
  ```

- **Example Response:**

  ```json
  {
      "message": "Payment intent created successfully",
      "payment_intent_id": "pi_1234567890",
      "amount": 1000
  }
  ```

### 2. Process Payment (To be implemented on the client)

- **URL:** `/process-payment`
- **Method:** `POST`
- **JSON Request Body:** 

  ```json
  {
      "token": "stripe_payment_token",
      "amount": 1000
  }
  ```

- **Example Response:**

  ```json
  {
      "message": "Payment processed successfully"
  }
  ```

## Workflow

1. To create a payment intent, make a POST request to `/create-payment-intent` with the desired payment amount in the request body.

2. The API will create a payment intent based on the provided data and respond with the intent ID.

3. To process the actual payment, you should implement the Stripe client on the client side (JavaScript) to collect card information and create a payment token. Then, you can make a POST request to `/process-payment` with the token and payment amount.

4. The API will confirm the payment intent using the received token and respond with a payment confirmation.

Be sure to follow best security practices when handling sensitive information and payment details.

---
This is a simple example of a Stripe payment API in Go. Remember to adjust and enhance this code as needed to meet the specific requirements of your project.

## Developed By:
**Francisco Inoque**
