HOMEWORK
========

An simple HTTP/JSON API written in GO.

## Installation
Install docker and cocker compose, checkout this repo.
```
docker-compose build api
docker-compose up api
docker-compose up database
```

## API Documentation 
### POST /driver
Create a driver
### POST /shipment
Create a shipment
### GET /driver/{id}
View a driver's offers
### GET /shipment/{id}
View a shipment's offer(s)
### PUT /offer/{id}
Accept or pass an offer

Test:
```
curl -H "Content-Type: application/json" -X POST -d '{"capacity": 100}' http://localhost:8080/driver
curl -H "Content-Type: application/json" -X POST -d '{"capacity": 50}' http://localhost:8080/driver
curl -H "Content-Type: application/json" -X POST -d '{"capacity": 50}' http://localhost:8080/shipment
curl -v http://localhost:8080/driver/1
curl -v http://localhost:8080/shipment/1
curl -H "Content-Type: application/json" -X PUT -d '{"status": "ACCEPT"}' http://localhost:8080/offer/1
```
