# Build the environment
# --------------------
FROM golang:1.17.1-alpine as builder

WORKDIR /myapp

RUN apk update && apk add --no-cache gcc musl-dev git bash

COPY . .

RUN go build -ldflags '-w -s' -a -o ./bin/app ./cmd/app

# Deployment environment
# ----------------------
FROM alpine

RUN apk update && apk add --no-cache bash

COPY --from=builder /myapp/bin/app /myapp/

EXPOSE 8000

CMD ["/myapp/app"]