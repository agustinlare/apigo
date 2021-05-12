FROM golang:1.16

ENV API_GO_PATH="/go/src/github.com/apigo"

USER 0
RUN mkdir -p ${API_GO_PATH}
COPY . ${API_GO_PATH}

ENV COUNTER_HIT_GOLANG  0

WORKDIR ${API_GO_PATH}

RUN go build
RUN chown -R 1001:0 ${API_GO_PATH}

USER 1001
EXPOSE 10000

CMD [ "./main" ]