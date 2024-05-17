# Prometheus metric validator

- This tool is a simple Prometheus metric validator. It reads a Prometheus metric from a URL and validates it against a
  set of rules.

## Getting started

### Build

```shell
make build
```

### Run

```shell
./prometheus-metric-validator --metric-url=http://localhost:9090/metrics
```

### Command line arguments
- `--metirc-url` - URL of the Prometheus metric
