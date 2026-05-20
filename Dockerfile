FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:3.19

WORKDIR /app

RUN adduser -D -g '' appuser

COPY --from=builder /app/main .

USER appuser

EXPOSE 8080

CMD ["./main"]