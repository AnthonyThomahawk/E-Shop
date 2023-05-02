# build stage
FROM --platform=${BUILDPLATFORM} golang:alpine as builder

# install delve debugger
# RUN go install github.com/go-delve/delve/cmd/dlv@latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /eshop-server 

# final stage 
FROM --platform=${BUILDPLATFORM} alpine as production

COPY --from=builder /eshop-server /

CMD [ "/eshop-server" ]
