# ApiGo

Simple example of an API in `golang` that counts every hit

# Build and push
```=bash
APIGO=quay.io/agustinlare/apigo
docker build -t $APIGO .
docker push $APIGO
```