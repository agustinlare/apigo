# ApiGo

A simple API in `golang` that counts every hit.

## Endpoints
+ `/health`: Status and message dependes on env.HEALTCHECK_STATUS
```
{
  "endpoint": "checklist",
  "ip": "172.17.0.1:57582",
  "counter": 1,
  "status": 200,
  "message": "true"
}
```
+ `/unhealthy`: Status 500
```
{
  "endpoint": "checklist",
  "ip": "172.17.0.1:57582",
  "counter": 1,
  "status": 500,
  "message": "cof...cof"
}
```
+ `/ping`: Status 200
```
{
  "endpoint": "checklist",
  "ip": "172.17.0.1:57582",
  "counter": 1,
  "status": 200,
  "message": "pong"
}
```
+ `/checklist`: Status 200, front 

### Front `/front`
+ IP Address, checks tcp conection to IP:PORT (port is 443 by default)
+ DNS, returns the associated ips
+ Mongodb, insert connection string to ping server and check if its reachable

## Build & run & push
```=bash
APIGO=quay.io/agustinlare/apigo

docker build -t $APIGO .
docker run -itd -p 8080:8080 $APIGO
docker push $APIGO
```
