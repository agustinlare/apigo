# ApiGo

A simple API in `golang` that counts every hit. Particularmente yo la use para estresar los pods del ingress controller de un cluster de Openshift 4.7

# Build and push
```=bash
APIGO=quay.io/agustinlare/apigo
docker build -t $APIGO .
docker push $APIGO
```
