FROM ubuntu:latest
LABEL authors="rianchoirulansyah"

ENTRYPOINT ["top", "-b"]
FROM golang:1.20-alphine AS build
WORKDIR /src
COPY go* ./
RUN apk add --no-cache git checkout -b
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -0 /bin/mau-pesan ./cmd/api

FROM alpine:3.18
RUN apk add --no-cache ca-certificates tzdata
COPY --from=build /bin/mau-pesan /app/mau-pesan
COPY .env /app/.env
EXPOSE 8080
CMD ["/app/mau-pesen"]