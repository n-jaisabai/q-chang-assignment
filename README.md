# Q-Chang assignment
## Project Description

This project is assignment of Q-Change interview.

* Backend is made with **Echo, Golang**.


## Run application with docker
1) Make sure docker and docker-compose is configured in your machine.
2) Run following command
    ```bash
    docker-compose build
    docker-compose up -d
    ```

## Run application without docker
To run it without using docker you need to have following installed:

  * You have to have **Golang 1.18** or above.
  * You have to have **postgres** up and running for database.
    
### Run backend project
1. install required packages:
    ```bash
    go mod tidy
    ```
2. Change [env](.env) file according to your database environment.

3. Run application. To run it in dev mode:
    ```bash
    go run main.go
    ```
4. If everything goes well and you have set your backend default running port as 8000

## Problem 1
  * Use ```GET /problem-1``` API to get answer of problem 1. For Example:
    ```bash
    curl -X 'GET' 'localhost:8000/problem-1'
    ```

## Problem 2
  * Use ```POST /problem-2``` API to get answer of problem 2 by given money of customer and product price. For Example:
    ```
    curl -X 'POST' \
      'http://localhost:8000/problem-2' \
      -H 'accept: application/json' \
      -H 'Content-Type: application/json' \
      -d '{
      "amount": 15,
      "product_price": 10
    }'
    ```
