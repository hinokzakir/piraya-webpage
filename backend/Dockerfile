FROM golang:1.24-alpine AS build

WORKDIR /app
COPY go.mod go.sum* ./
COPY main.go handlers.go ./
RUN go mod tidy && CGO_ENABLED=0 GOOS=linux go build -o server .

FROM alpine:3.21

WORKDIR /app
COPY --from=build /app/server /app/server
RUN mkdir -p /data

ENV PORT=8080
ENV DB_PATH=/data/piraya.db

EXPOSE 8080
CMD ["/app/server"]
