FROM golang:1.21

WORKDIR /app

ADD . .

RUN go mod download
RUN GOOS=linux GOARCH=amd64 go build -o /api ./

EXPOSE 8080

CMD ["/api"]
