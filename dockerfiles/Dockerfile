FROM golang:latest as builder

WORKDIR /app

COPY . .

RUN go mod download && \
    CGO_ENABLED=0 go build -o ./server

FROM scratch
COPY --from=builder /app/server /server
ENTRYPOINT ["/server"]