FROM golang:1.18 as builder

RUN mkdir /artifact

WORKDIR /workspace

COPY . .

RUN go build -o /artifact/app ./cmd/server

FROM gcr.io/distroless/base-debian10

COPY --from=builder /artifact/app /app

RUN apt-get -y update && \
    apt-get -y install jq
ENTRYPOINT ["/entrypoint.sh"]

CMD [ "/app" ]