FROM golang:1.16

ENV API_GO_PATH="/go/src/github.com/apigo"

RUN mkdir -p ${API_GO_PATH}
COPY . ${API_GO_PATH}

EXPOSE 10000

WORKDIR ${API_GO_PATH}

ENTRYPOINT [ "go", "run", "main.go" ]