FROM golang:1.25.3 AS subscription-service

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main ./cmd/app

EXPOSE 9091

CMD [ "./main" ]