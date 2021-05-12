APIGO=quay-enterprise-quay-quay-enterprise.apps.ocppaz0.ar.bsch/a309788/apigo:1
docker build -t $APIGO .
docker push $APIGO