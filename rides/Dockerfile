FROM golang:1.23-alpine3.20 AS deps

RUN apk update && apk add --no-cache git ca-certificates

FROM deps AS pkgs

WORKDIR /build/app

COPY rides/go.mod ./rides/
COPY rides/go.sum ./rides/
COPY shared/go.mod ./shared/
COPY shared/go.sum ./shared/
COPY proto/go.mod ./proto/
COPY proto/go.sum ./proto/

RUN cd rides && go mod download

FROM pkgs AS builder

COPY rides rides
COPY shared shared
COPY proto proto

RUN cd rides && go build -o /app/app .

FROM scratch AS runner

WORKDIR /app

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/app ./

ENTRYPOINT ["/app/app"]
