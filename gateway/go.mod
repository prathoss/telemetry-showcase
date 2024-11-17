module github.com/prathoss/telemetry_showcase/gateway

go 1.22.6

replace (
	github.com/prathoss/telemetry_showcase/proto => ../proto
	github.com/prathoss/telemetry_showcase/shared => ../shared
)

require (
	github.com/99designs/gqlgen v0.17.49
	github.com/google/uuid v1.6.0
	github.com/prathoss/telemetry_showcase/proto v0.0.0-00010101000000-000000000000
	github.com/prathoss/telemetry_showcase/shared v0.0.0-00010101000000-000000000000
	github.com/ravilushqa/otelgqlgen v0.16.0
	github.com/vektah/gqlparser/v2 v2.5.16
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.54.0
	google.golang.org/genproto v0.0.0-20240827150818-7e3bb234dfed
	google.golang.org/grpc v1.66.0
)

require (
	github.com/agnivade/levenshtein v1.1.1 // indirect
	github.com/cenkalti/backoff/v4 v4.3.0 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.4 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-ole/go-ole v1.3.0 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.22.0 // indirect
	github.com/hashicorp/golang-lru/v2 v2.0.7 // indirect
	github.com/lufia/plan9stats v0.0.0-20240819163618-b1d8f4d146e7 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/power-devops/perfstat v0.0.0-20240221224432-82ca36839d55 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/shirou/gopsutil/v4 v4.24.7 // indirect
	github.com/shoenig/go-m1cpu v0.1.6 // indirect
	github.com/sosodev/duration v1.3.1 // indirect
	github.com/tklauser/go-sysconf v0.3.14 // indirect
	github.com/tklauser/numcpus v0.8.0 // indirect
	github.com/urfave/cli/v2 v2.27.2 // indirect
	github.com/xrash/smetrics v0.0.0-20240312152122-5f08fbb34913 // indirect
	github.com/yusufpapurcu/wmi v1.2.4 // indirect
	go.opentelemetry.io/contrib v1.27.0 // indirect
	go.opentelemetry.io/contrib/bridges/otelslog v0.4.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/host v0.54.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/runtime v0.54.0 // indirect
	go.opentelemetry.io/otel v1.29.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploggrpc v0.5.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc v1.29.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.29.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.29.0 // indirect
	go.opentelemetry.io/otel/log v0.5.0 // indirect
	go.opentelemetry.io/otel/metric v1.29.0 // indirect
	go.opentelemetry.io/otel/sdk v1.29.0 // indirect
	go.opentelemetry.io/otel/sdk/log v0.5.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v1.29.0 // indirect
	go.opentelemetry.io/otel/trace v1.29.0 // indirect
	go.opentelemetry.io/proto/otlp v1.3.1 // indirect
	golang.org/x/mod v0.18.0 // indirect
	golang.org/x/net v0.28.0 // indirect
	golang.org/x/sync v0.8.0 // indirect
	golang.org/x/sys v0.24.0 // indirect
	golang.org/x/text v0.17.0 // indirect
	golang.org/x/tools v0.22.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20240823204242-4ba0660f739c // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240827150818-7e3bb234dfed // indirect
	google.golang.org/protobuf v1.34.2 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
