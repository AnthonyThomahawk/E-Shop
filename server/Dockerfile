FROM golang:1.20.2-alpine3.17 

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /eshop-server 

CMD [ "/eshop-server" ]
