FROM golang:1.18-alpine

WORKDIR /app

COPY . ./
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o ./out/contactbook-project .
EXPOSE 8080

CMD ["./out/contactbook-project"]