FROM golang:alpine3.17 as deps

RUN apk update && apk add --no-cache git ca-certificates

FROM deps as pkgs

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

FROM pkgs as builder

COPY . .

RUN go build -o /build/app .

FROM scratch as runner

WORKDIR /app

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /build/app ./

ENTRYPOINT ["/app/app"]
