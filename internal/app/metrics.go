package app

import (
	"github.com/prometheus/client_golang/prometheus"
	"paint/pkg/def"
)

var (
	Metric def.Metrics
)

func InitMetrics(reg *prometheus.Registry) {
	Metric = def.NewMetrics(reg)
}
