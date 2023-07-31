FROM golang:latest

WORKDIR /app

COPY . /app
RUN go mod download

RUN go build -o scrim-bot .

EXPOSE 8080

CMD ["/app/scrim-bot"]