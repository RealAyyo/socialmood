FROM golang:latest
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download
RUN go mod tidy

COPY ./ ./
RUN go get ./...
RUN go build -o socialmood ./cmd/social-mood

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

COPY ./migrations ./migrations

ENTRYPOINT /bin/sh -c "goose -dir ./migrations postgres \"postgres://$DB_USER:$DB_PASSWORD@$DB_HOST:$DB_PORT/$DB_NAME\" up && ./socialmood --config=configs/config.yaml"