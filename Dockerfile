FROM golang:1.21

WORKDIR /app

ADD . .

RUN go mod download
RUN COMMIT=$(git log -1 --format='%H') && \
    GOOS=linux GOARCH=amd64 go build -ldflags="-X github.com/cosmos/cosmos-sdk/version.Commit=${COMMIT}" -o /api ./

EXPOSE 8080

CMD ["/api"]
