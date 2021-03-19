package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"paint/pkg/def"
	"runtime"
)

func initMetrics(reg *prometheus.Registry, namespace string) {
	reg.MustRegister(prometheus.NewProcessCollector(prometheus.ProcessCollectorOpts{}))
	reg.MustRegister(prometheus.NewGoCollector())

	version := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "build_info",
			Help:      "A metric with a constant '1' value labeled by build-time details.",
		},
		[]string{"version", "goversion"},
	)
	reg.MustRegister(version)

	version.With(prometheus.Labels{
		"version":   def.Version(),
		"goversion": runtime.Version(),
	}).Set(1)
}
