# MicroServices-Ecommerce

MicroServices-Ecommerce is an e-commerce platform built on a microservices architecture. It leverages gRPC for communication between services, Google OAuth 2.0 for user authentication, and Stripe as the payment gateway.

## Features

- **gRPC Integration**: MicroServices-Ecommerce utilizes gRPC (Google Remote Procedure Call) for communication between its microservices. gRPC is a high-performance, open-source RPC framework developed by Google.

- **Google OAuth 2.0 Authentication**: Users can securely sign in to MicroServices-Ecommerce using their Google accounts. The platform integrates Google OAuth 2.0 for user authentication, ensuring a seamless and secure login experience.

- **Stripe Payment Gateway**: MicroServices-Ecommerce supports Stripe as the payment gateway for processing payments. With Stripe integration, users can make purchases securely using various payment methods, including credit/debit cards and digital wallets.

## Getting Started

To get started with MicroServices-Ecommerce, follow these steps:

1. Clone the repository: `git clone https://github.com/Vajid-Hussain/MicroServices-Ecommerce.git`
2. Install dependencies for each microservice: Navigate to the directory of each microservice and run `npm install` or `go mod tidy` to install dependencies.
3. Configure environment variables: Update the `.env` files in each microservice directory with your credentials and configurations.
4. Start each microservice: Run each microservice individually by navigating to its directory and executing the appropriate command (e.g., `npm start` or `go run main.go`).
5. Open your browser and navigate to the appropriate URL to access MicroServices-Ecommerce.

## Contributing

Contributions to MicroServices-Ecommerce are welcome! If you have any ideas for new features, bug fixes, or improvements, feel free to open an issue or submit a pull request.

