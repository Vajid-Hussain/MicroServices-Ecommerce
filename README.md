
# MicroServices-Ecommerce

A scalable e-commerce platform built using a microservices architecture. Leverages gRPC for communication between services, Google OAuth 2.0 for user authentication, and Stripe as the payment gateway.

## Features

* **Secure Authentication:** User registration and login protected by Google OAuth 2.0. 
* **Product Management:** Comprehensive product details, inventory, and pricing management.
* **Order Processing:** Facilitates streamlined order creation and payment processing using Stripe.
* **Scalable Architecture:** Microservices architecture designed for flexibility and growth.

## Architecture

* **API Gateway:** Handles client requests, routing, and Google OAuth 2.0 authentication.
* **User Authentication Service:** Manages user data and credentials.
* **Product Service:** Central repository for product-related information.
* **Order Service:** Handles order creation, payment processing (Stripe), and order lifecycle.

## Technologies

* **gRPC:** For efficient inter-service communication.
* **Google OAuth 2.0:** Secure user authentication.
* **Stripe:** Payment processing integration.
* **Languages:** (Golang)
* **Databases:** (postgresSql)

## Getting Started

1. **Prerequisites:** 
2. **Clone:** `git clone https://github.com/Vajid-Hussain/MicroServices-Ecommerce.git`
3. **Run:** make run


