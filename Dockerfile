FROM golang:1.15-alpine AS build_base

RUN apk add --no-cache git
WORKDIR /tmp/subscription-service

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o ./out/subscription-service .

FROM scratch
LABEL maintainer="Julian Klaiber"

COPY --from=build_base /tmp/subscription-service/out/subscription-service /usr/bin/subscription-service

ENTRYPOINT [ "/usr/bin/subscription-service" ]