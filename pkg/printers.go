package pkg

import (
	"github.com/phin1x/go-ipp"
	"github.com/prometheus/client_golang/prometheus"
)

func (e *Exporter) printerMetrics(ch chan<- prometheus.Metric) error {
	printers, err := e.client.GetPrinters([]string{"printer-state"})
	if err != nil {
		e.log.Error(err, "failed to fetch completed jobs")
		return err
	}

	states := make(map[ipp.PrinterState]int)
	states[ipp.PrinterStateIdle] = 0
	states[ipp.PrinterStateProcessing] = 0
	states[ipp.PrinterStateStopped] = 0

	for _, attr := range printers {
		states[ipp.PrinterState(attr["printer-state"][0].Value.(int))]++
	}

	ch <- prometheus.MustNewConstMetric(e.printersTotal, prometheus.GaugeValue, float64(len(printers)))
	ch <- prometheus.MustNewConstMetric(e.printerStateTotal, prometheus.GaugeValue, float64(states[ipp.PrinterStateIdle]), "idle")
	ch <- prometheus.MustNewConstMetric(e.printerStateTotal, prometheus.GaugeValue, float64(states[ipp.PrinterStateProcessing]), "processing")
	ch <- prometheus.MustNewConstMetric(e.printerStateTotal, prometheus.GaugeValue, float64(states[ipp.PrinterStateStopped]), "stopped")

	return nil
}
