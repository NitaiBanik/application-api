#!/bin/bash

lsof -ti:8080,8081,8082 | xargs kill -9
sleep 2

PORT=8080 go run main.go &
PORT=8081 go run main.go &
PORT=8082 go run main.go &

sleep 3

curl -X POST http://localhost:8080/testapi -H "Content-Type: application/json" -d '{"test": 1}'
curl -X POST http://localhost:8081/testapi -H "Content-Type: application/json" -d '{"test": 2}'
curl -X POST http://localhost:8082/testapi -H "Content-Type: application/json" -d '{"test": 3}'

wait
