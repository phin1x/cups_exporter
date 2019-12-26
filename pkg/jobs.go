package pkg

import (
	"github.com/phin1x/go-ipp"
	"github.com/prometheus/client_golang/prometheus"
)

func (e *Exporter) jobsMetrics(ch chan<- prometheus.Metric) error {
	jobs, err := e.client.GetJobs("", "", ipp.JobStateFilterNotCompleted, false, 0, 0, []string{})
	if err != nil {
		e.log.Error(err, "failed to fetch completed jobs")
		return err
	}

	activeJobs := len(jobs)

	jobs, err = e.client.GetJobs("", "", ipp.JobStateFilterAll, false, e.lastJobId, 0, []string{})
	if err != nil {
		e.log.Error(err, "failed to fetch all jobs")
		return err
	}

	lastJobId := getLastJobId(jobs)
	if lastJobId > e.lastJobId {
		e.lastJobId = lastJobId
	}

	ch <- prometheus.MustNewConstMetric(e.jobsTotal, prometheus.CounterValue, float64(lastJobId))
	ch <- prometheus.MustNewConstMetric(e.jobsActiveTotal, prometheus.GaugeValue, float64(activeJobs))

	return nil
}

/*
	returns the last job id, the last completed job id and the current active jobs
*/
func getLastJobId(m map[int]ipp.Attributes) int {
	lastJobId := 0

	for k := range m {
		if k > lastJobId {
			lastJobId = k
		}
	}

	return lastJobId
}
