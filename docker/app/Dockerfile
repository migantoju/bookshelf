# Build the environment
# --------------------
FROM golang:1.17.1-alpine as builder

WORKDIR /myapp

RUN apk update && apk add --no-cache gcc musl-dev git bash

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -ldflags '-w -s' -a -o ./bin/app ./cmd/app \
    && go build -ldflags '-w -s' -a -o ./bin/migrate ./cmd/migrate

# Deployment environment
# ----------------------
FROM alpine

RUN apk update && apk add --no-cache bash

COPY --from=builder /myapp/bin/app /myapp/
COPY --from=builder /myapp/bin/migrate /myapp/
COPY --from=builder /myapp/migrations /myapp/migrations

EXPOSE 8000

CMD ["/myapp/app"]