# ApiGo

A simple API in `golang` that counts every hit.

## Endpoints
+ `/health`: Status 200, `{"response": "ok"}`
+ `/unhealthy`: Status 500, `{"response": "error"}`
+ `/ping`: Status 200, `{"ping": "pong"}`

## Build & run & push
```=bash
APIGO=quay.io/agustinlare/apigo

docker build -t $APIGO .
docker run -itd -p 8080:8080 $APIGO
docker push $APIGO
```
