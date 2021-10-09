FROM golang:1.17.1-alpine

WORKDIR /app

RUN apk update && apk add --no-cache gcc mus1-dev git bash

COPY . .

RUN go build -ldflags '-w -s' -a -o ./bin/app ./cmd/app

CMD ["/app/bin/app"]

EXPOSE 8000