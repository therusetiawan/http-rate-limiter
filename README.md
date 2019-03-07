# HTTP Rate Limiter

### What is Rate Limiter

Rate limiter is a mechanism to control the amount of incoming request to server. For example we are
only accept 10 requests/second. If the number of requests you make exceeds that limit,
then an error will be triggered. One of the benefit implement rate limiter is to prevent
DDoS attack. 

### Requirements
* Golang 1.9 or later
* Redis server

### How to Run

1. Install all project dependencies
```
make install-dep
```
2. Start docker service using docker-compose
```
docker-compose up -d
```
3. Run http rate limiter
```
go run main.go -listen-port 9000 rate-limit=10 -redis-url localhost:6378
```

### How to Test

The easiest way is using ![Postman Collection Run] (https://learning.getpostman.com/docs/postman/collection_runs/starting_a_collection_run/) to make flood requests.

```
curl -X GET \
  'http://localhost:9000/api/v1/ping?user_token=usertoken' \
  -H 'Content-Type: application/json' \
  -H 'Postman-Token: d7439965-b874-4c39-bf2c-875f9948e012' \
  -H 'cache-control: no-cache' \
  -H 'content-type: multipart/form-data; boundary=----WebKitFormBoundary7MA4YWxkTrZu0gW' \
  -F author=Heru \
  -F title=RateLimiter
```

## Built With

* [gin](https://github.com/gin-gonic/gin) - HTTP web framework
* [go-redis](https://github.com/go-redis/redis) - Redis client for Golang
