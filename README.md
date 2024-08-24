# Cript Core Customer

![Go](https://img.shields.io/badge/Go-1.21.5-blue)
![License](https://img.shields.io/badge/License-MIT-green)

## Description

This project is part of my project that simulate a bitcoin broker.
more details [here](https://andreleao1.github.io/criptotrade.github.io/).

This project is write in Go and your propolse of this service is manage all customer expirience in the broker like register a new customer, create wallet and etc. 

## Requiments
- Go
- Docker

## Features

- Register and manager Customers
- Login
- Create and manage customers's wallets

## Whats is used here

- Gin [see documentation](https://github.com/gin-gonic/gin)
- Sqlx [see documentation](https://github.com/jmoiron/sqlx)
- uuid [see documentation](https://github.com/google/uuid)
- migrate [see documentation](https://github.com/golang-migrate/migrate)

## Intallation
To install and run the project, follow these steps:

1. **Clone the repository:**
    ```sh
    git clone https://github.com/andreleao1/cript-core-customer.git
    cd cript-core-customer
    ```

2. **Install dependencies:**
    ```sh
    go mod tidy
    ```

3. **Run docker compose file**

    ```sh
    cd src
    docker-compose up -d
    ```

3. **Build the project:**
    ```sh
    go run .
    ```
