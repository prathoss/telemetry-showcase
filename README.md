# Telemetry showcase

```mermaid
flowchart LR
    gateway --> users
    gateway --> bikes
    gateway --> rides
    users --> redis
    redis --> pg
    bikes --> redis
    bikes --> pg
    rides --> kafka
    rides --> bikes
    rides --> pg
    kafka <--> invoices
```

Simple setup of services:
- gql is GraphQL system gateway
- users is gRPC service handling the user entities
- bikes is gRPC service handling the bike entities
- invoices handles processing and sending invoices

these services are instrumented to create logs, metrics and traces through OTEL.

Project also contains docker compose files to spin up previously mentioned services and also stacks to
analyse and visualise telemetry data.

## Requirements

Collectors:
- [x] Grafana showcase
- [ ] ELK stack showcase
- [ ] AWS showcase
- [x] DataDog showcase
- [ ] Maybe more

Metrics:
- [x] Runtime
- [x] Host
- [x] gRPC
- [x] GraphQL
- [ ] BlackBox exporter (alloy)
- [x] Redis exporter (alloy)
- [x] Kafka exporter (alloy)
- [x] Postgres exporter (alloy)

Tracing:
- [x] gRPC
- [x] GraphQL
- [x] Custom (simple)
- [x] Redis
- [x] Gorm
- [x] Kafka
- [x] Linked traces
- [ ] Events
- [x] Error in span

### Sampling

It is possible to use a rate-based collection in SDK, but it is quite simple
and may skip information which is important for situation assessment, this rating may discard, for example,
trace with error.
Advanced sampling should be configured at agent or collector which will have more information
about the whole trace.
For more information, take a look at [head vs tail sampling](https://opentelemetry.io/docs/concepts/sampling/#head-sampling).

## What to look forward to

- [profiling announcement](https://opentelemetry.io/blog/2024/profiling/)

## Configuration

- [Exporter](https://opentelemetry.io/docs/specs/otel/protocol/exporter/)
- [SDK](https://opentelemetry.io/docs/specs/otel/configuration/sdk-environment-variables/#general-sdk-configuration)

### Datadog exporter

Go to file `.docker/datadog/otel-collector.yml:51` and substitute `{{key}}` with your datadog api key.

