FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

FROM alpine
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .
COPY start.sh .
COPY wait-for.sh .
COPY db/migration ./db/migration

EXPOSE 8082
CMD ["/app/main"]
ENTRYPOINT [ "/app/start.sh" ]
