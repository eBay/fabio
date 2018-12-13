package prometheus

import (
	"github.com/fabiolb/fabio/metrics4"
	"sync"
	"time"

	"github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

type Provider struct {
	counters   map[string]metrics4.Counter
	gauges     map[string]metrics4.Gauge
	timers     map[string]metrics4.Timer
	mutex      sync.Mutex
}

func NewProvider() *Provider {
	return &Provider{
		counters: make(map[string]metrics4.Counter),
		gauges: make(map[string]metrics4.Gauge),
		timers: make(map[string]metrics4.Timer),
	}
}

func (p *Provider) NewCounter(name string, labels ... string) metrics4.Counter {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	if p.counters[name] == nil {
		p.counters[name] = prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: metrics4.FabioNamespace,
			Subsystem: "",
			Name:      name,
			Help:      "",
		}, labels)
	}

	return p.counters[name]
}

func (p *Provider) NewGauge(name string, labels ... string) metrics4.Gauge {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	if p.gauges[name] == nil {
		p.gauges[name] = prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: metrics4.FabioNamespace,
			Subsystem: "",
			Name:      name,
			Help:      "",
		}, labels)
	}

	return p.gauges[name]
}

//func (p *Provider) NewHistogram(name string, labels ... string) metrics4.Histogram {
//	p.mutex.Lock()
//	defer p.mutex.Unlock()
//	if p.histograms[name] == nil {
//		p.histograms[name] = prometheus.NewHistogramFrom(stdprometheus.HistogramOpts{
//			Namespace: metrics4.FabioNamespace,
//			Subsystem: "",
//			Name:      name,
//			Help:      "",
//			// TODO: Look on 'Buckets'
//		}, labels)
//	}
//
//	return p.histograms[name]
//}

func (p *Provider) NewTimer(name string, labels ... string) metrics4.Timer {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	if p.timers[name] == nil {
		h := prometheus.NewHistogramFrom(stdprometheus.HistogramOpts{
			Namespace: metrics4.FabioNamespace,
			Name:      name,
		}, labels)

		p.timers[name] = metrics4.NewTimerStruct(h, time.Now())
	}

	return p.timers[name]
}
