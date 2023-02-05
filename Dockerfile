FROM golang:1.19.5-alpine
WORKDIR /app
ADD . .
RUN go build -o app cmd/main.go 
CMD ["./app"]