# BASIC BANKING TRANSACTION IN GO

This project is an API to simulate banking transactions in Golang.

## Functionalities

POST /accounts => Create a new account with $1000,00 on balance

GET /accounts => Retrieve all accounts

GET /accounts/:id => Retrieve a specific account

GET /accounts/:id/balance => Retrieve the balance of an account

POST /transfers => Make a transaction to one logged in account to another account

GET /transfers => Retrieve all transactions of the logged in account

POST /login => Retrieve a JWT token to be used transactions api

POST /logout => Destroy the token to avoid attacks.

## Pre requisites

* Docker
* Docker compose

## Executing

```
docker-compose build 
docker-compose up -d
```

It will create container that run the necessary services.

The services will be running on port 8080
