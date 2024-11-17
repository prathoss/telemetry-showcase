FROM golang:1.23-alpine3.20 AS deps

RUN apk update && apk add --no-cache git ca-certificates

FROM deps AS pkgs

WORKDIR /build/app

COPY users/go.mod ./users/
COPY users/go.sum ./users/
COPY shared/go.mod ./shared/
COPY shared/go.sum ./shared/
COPY proto/go.mod ./proto/
COPY proto/go.sum ./proto/

RUN cd users && go mod download

FROM pkgs AS builder

COPY users users
COPY shared shared
COPY proto proto

RUN cd users && go build -o /app/app .

FROM scratch AS runner

WORKDIR /app

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/app ./

ENTRYPOINT ["/app/app"]
