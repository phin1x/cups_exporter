package main

import (
	"flag"
	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"github.com/phin1x/cups-exporter/internal"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"net/http"
	"os"
)

type options struct {
	Address     string
	MetricsPath string

	CupsUri string
}

func main() {
	log := getLogger()

	opts := options{}
	flag.StringVar(&opts.Address, "web.listen-address", ":9628", "address on which to expose metrics and web interface")
	flag.StringVar(&opts.MetricsPath, "web.telemetry-path", "/metrics", "path under which to expose metrics")
	flag.StringVar(&opts.CupsUri, "cups.uri", "https://localhost:631", "uri under with the cups server is available, including username and password it required")
	flag.Parse()

	log.Info("starting cups exporter")
	exporter, err := internal.NewExporter(opts.CupsUri, log)
	if err != nil {
		log.Error(err, "failed to create the exporter")
		os.Exit(1)
	}
	prometheus.MustRegister(exporter)

	http.Handle(opts.MetricsPath, promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
			<head><title>Cups Exporter</title></head>
			<body>
			<h1>Cups Exporter</h1>
			<p><a href="` + opts.MetricsPath + `">Metrics</a></p>
			</body>
			</html>`))
	})

	log.Info("listening on " + opts.Address)
	log.Error(http.ListenAndServe(opts.Address, nil), "failed to start the http server")
}

func getLogger() logr.Logger {
	zapLog, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	return zapr.NewLogger(zapLog)
}
