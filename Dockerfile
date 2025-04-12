# Stage builder
FROM golang:1.23.8-alpine as builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o dietcalc ./cmd/DietCalc

# Stage base
FROM alpine:latest as base
WORKDIR /root
COPY --from=builder /app/dietcalc .
COPY --from=builder /app/internal/storage/postgres/migrations ./migrations
RUN apk --no-cache add ca-certificates

# Stage dev
FROM base as dev
EXPOSE 3001
CMD ["./dietcalc"]

# Stage prod
FROM base as prod
EXPOSE 3000
CMD ["./dietcalc"]
