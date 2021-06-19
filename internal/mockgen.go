package internal

//go:generate mockgen -destination=./mocks/flusher_mock.go -package=mocks github.com/ozoncp/ocp-check-api/internal/flusher CheckFlusher,TestFlusher
//go:generate mockgen -destination=./mocks/repo_mock.go -package=mocks github.com/ozoncp/ocp-check-api/internal/repo CheckRepo,TestRepo
//go:generate mockgen -destination=./mocks/alarmer_mock.go -package=mocks github.com/ozoncp/ocp-check-api/internal/alarmer Alarmer
//go:generate mockgen -destination=./mocks/producer_mock.go -package=mocks github.com/ozoncp/ocp-check-api/internal/producer Producer
//go:generate mockgen -destination=./mocks/prometheus_mock.go -package=mocks github.com/ozoncp/ocp-check-api/internal/prometheus Prometheus
