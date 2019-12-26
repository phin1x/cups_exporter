# Prometheus CUPS Exporter

[![Version](https://img.shields.io/github/release-pre/phin1x/cups_exporter.svg)](https://github.com/phin1x/cups_exporter/releases/tag/v1.0.0)
[![Licence](https://img.shields.io/github/license/phin1x/cups_exporter.svg)](https://github.com/phin1x/cups_exporter/blob/master/LICENSE)

Prometheus exporter for CUPS server

# Build

```bash
go build -o cups_exporter main.go
```

# Running

By default the cups_exporter serves on port `0.0.0.0:9628` at `/metrics`. The cups server is specified by the `cups.uri` flag (default: `https://localhost:631`).

Examples:
```bash
./cups_exporter # use defaults
./cups_exporter -cups.uri https://exporter:prometheus@mycups.foo.bar:631 # scrape remote server with basic auth
```

# Metrics

| Metric | Meaning | Labels |
| ------ | ------- | ------ |
| up | Was the last scrape of cups successful | |
| cups_job_active_total | Number of current print jobs | |
| cups_job_total | Total number of print jobs | |
| cups_printer_state_total | Number of printers per state | state |
| cups_printer_total | Number of available printers | |
| cups_scrape_duration_seconds |  Duration of the last scrape in seconds | |

# Licence

Apache Licence Version 2.0
