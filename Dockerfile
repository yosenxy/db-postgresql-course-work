FROM golang:1.25.4
LABEL authors="yosenxy"

WORKDIR /db

COPY . .

RUN go mod download

RUN go mod tidy

RUN go build -o app ./cmd/service/main.go

EXPOSE 8080

CMD ["/app"]