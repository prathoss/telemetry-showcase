rebuild-grafana:
	docker compose -f compose.grafana.yml up -d --build

rebuild-datadog:
	docker compose -f compose.datadog.yml up -d --build

alloy-fmt:
	docker run --rm --workdir /src -v "$$(pwd)/.docker:/docker" grafana/alloy:v1.3.1 fmt --write /docker/grafana/config.alloy
