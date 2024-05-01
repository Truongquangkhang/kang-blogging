FROM golang:1.18-alpine

ARG SERVICE

WORKDIR sql

COPY go.mod ./
COPY go.sum ./
RUN go mod download

RUN mkdir -p $SERVICE/migrations
COPY main.go ./
COPY $SERVICE/migrations/*.sql ./$SERVICE/migrations/

RUN go build -o /sql .

CMD ["/sql", "migrate", "up"]
