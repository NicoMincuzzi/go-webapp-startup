FROM golang:1.16.2-alpine3.13 AS builder

WORKDIR /app
COPY go.mod go.sum /app/

## Add this go mod download command to pull in any dependencies
RUN go mod download

COPY ./cmd ./cmd

# Unit tests
#RUN CGO_ENABLED=0 go test -v

RUN go build -o /build/ cmd/*.go

FROM alpine:3.13

COPY --from=builder /build /bin/cmd
WORKDIR /app

EXPOSE 3030

CMD ["./main"]