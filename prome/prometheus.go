package prome

import (
	"github.com/prometheus/client_golang/prometheus"
	"time"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"log"
	"github.com/x1957/dockermonitor/agent"
)

type Prome struct {
	// no concurrency
	counters map[string]prometheus.Gauge
}

func NewPrometheus() agent.Agent {
	return Prome{
		counters: make(map[string]prometheus.Gauge),
	}
}

func (p Prome) RecordGauge(name string, gaugeValue int64) {
	gauge, ok := p.counters[name]
	if !ok {
		// registe the counter
		gauge = prometheus.NewGauge(prometheus.GaugeOpts{
			Name: name,
			Help: name,
		})
		prometheus.MustRegister(gauge)
		p.counters[name] = gauge
	}
	gauge.Set(float64(gaugeValue))
}

func (p Prome) RecordDuration(name string, duration time.Duration) {
	gauge, ok := p.counters[name]
	if !ok {
		// registe the counter
		gauge = prometheus.NewGauge(prometheus.GaugeOpts{
			Name: name,
			Help: name,
		})
		prometheus.MustRegister(gauge)
		p.counters[name] = gauge
	}
	gauge.Set(float64(duration))
}

func (p Prome) Run() {
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		log.Fatal(http.ListenAndServe(":1957", nil))
	}()
}