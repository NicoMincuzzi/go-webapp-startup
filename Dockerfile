FROM golang:1.16.2-alpine3.13 AS builder

WORKDIR /app

# COPY go.mod, go.sum and download the dependencies
COPY go.* ./
RUN go mod download

COPY ./cmd/*.go ./

RUN go build -o /go-webapp

FROM alpine:3.13

COPY --from=builder /go-webapp /go-webapp

EXPOSE 3030

CMD ["/go-webapp"]