package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/rs/zerolog"
)

type Prometheus interface {
	IncCreateCheck(status string)
	IncUpdateCheck(status string)
	IncDeleteCheck(status string)
	IncCreateTest(status string)
	IncUpdateTest(status string)
	IncDeleteTest(status string)
}

type prometheusApi struct {
	createCounter  *prometheus.CounterVec
	updateCounter  *prometheus.CounterVec
	deleteCounter  *prometheus.CounterVec
	createCounterT *prometheus.CounterVec
	updateCounterT *prometheus.CounterVec
	deleteCounterT *prometheus.CounterVec
}

func NewPrometheus(log zerolog.Logger) Prometheus {
	api := &prometheusApi{
		createCounter: promauto.NewCounterVec(
			prometheus.CounterOpts{
				Name: "ocp_check_create_counter",
				Help: "Number of created checks.",
			},
			[]string{"status"},
		),

		updateCounter: promauto.NewCounterVec(
			prometheus.CounterOpts{
				Name: "ocp_check_update_counter",
				Help: "Number of updated checks.",
			},
			[]string{"status"},
		),

		deleteCounter: promauto.NewCounterVec(
			prometheus.CounterOpts{
				Name: "ocp_check_delete_counter",
				Help: "Number of deleted checks.",
			},
			[]string{"status"},
		),

		createCounterT: promauto.NewCounterVec(
			prometheus.CounterOpts{
				Name: "ocp_test_create_counter",
				Help: "Number of created tests.",
			},
			[]string{"status"},
		),

		updateCounterT: promauto.NewCounterVec(
			prometheus.CounterOpts{
				Name: "ocp_test_update_counter",
				Help: "Number of updated tests.",
			},
			[]string{"status"},
		),

		deleteCounterT: promauto.NewCounterVec(
			prometheus.CounterOpts{
				Name: "ocp_test_delete_counter",
				Help: "Number of deleted tests.",
			},
			[]string{"status"},
		)}

	return api
}

func (p *prometheusApi) IncCreateCheck(status string) {
	p.createCounter.WithLabelValues(status).Inc()
}

func (p *prometheusApi) IncUpdateCheck(status string) {
	p.updateCounter.WithLabelValues(status).Inc()
}

func (p *prometheusApi) IncDeleteCheck(status string) {
	p.deleteCounter.WithLabelValues(status).Inc()
}

func (p *prometheusApi) IncCreateTest(status string) {
	p.createCounterT.WithLabelValues(status).Inc()
}

func (p *prometheusApi) IncUpdateTest(status string) {
	p.updateCounterT.WithLabelValues(status).Inc()
}

func (p *prometheusApi) IncDeleteTest(status string) {
	p.deleteCounterT.WithLabelValues(status).Inc()
}
