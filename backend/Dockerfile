FROM golang:1.23-alpine3.21

RUN apk add curl
WORKDIR /app
COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /main

EXPOSE 8080

HEALTHCHECK --start-period=5s CMD curl --fail localhost:8080/api/health || exit 1

CMD [ "/main" ]
