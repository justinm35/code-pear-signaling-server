FROM golang:1.23

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /go-signaling-server

EXPOSE 4000

CMD ["/go-signaling-server"]
