package dal

import (
	"github.com/prometheus/client_golang/prometheus"
	"paint/internal/app"
	"paint/pkg/repo"
)

var metric repo.Metrics

func InitMetrics(reg *prometheus.Registry, namespace string) {
	const subsystem = "dal_sql"

	metric = repo.NewMetrics(reg, namespace, subsystem, new(app.Repo))
}
