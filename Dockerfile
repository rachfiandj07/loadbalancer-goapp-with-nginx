FROM golang:1.15-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /loadbalancer-go-nginx

EXPOSE 8080

ENTRYPOINT [ "/loadbalancer-go-nginx" ]