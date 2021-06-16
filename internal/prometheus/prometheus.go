package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/rs/zerolog"
)

type Prometheus interface {
	IncCreateCheck(status string)
	IncUpdateCheck(status string)
	IncDeleteCheck(status string)
}

type prometheusApi struct {
	createCounter *prometheus.CounterVec
	updateCounter *prometheus.CounterVec
	deleteCounter *prometheus.CounterVec
}

func NewPrometheus(log zerolog.Logger) Prometheus {
	api := &prometheusApi{
		createCounter: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "ocp_check_create_counter",
				Help: "Number of created checks.",
			},
			[]string{"status"},
		),

		updateCounter: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "ocp_check_update_counter",
				Help: "Number of updated checks.",
			},
			[]string{"status"},
		),

		deleteCounter: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "ocp_check_delete_counter",
				Help: "Number of deleted checks.",
			},
			[]string{"status"},
		)}

	if err := prometheus.Register(api.createCounter); err != nil {
		log.Error().Err(err).Msg("")
	}
	if err := prometheus.Register(api.updateCounter); err != nil {
		log.Error().Err(err).Msg("")
	}
	if err := prometheus.Register(api.deleteCounter); err != nil {
		log.Error().Err(err).Msg("")
	}

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
