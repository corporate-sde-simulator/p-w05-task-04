package metrics

// MetricsCollector — registers and tracks metric counters and gauges.
// This module is COMPLETE and WORKING. Do NOT modify this file.
// Your task is to implement the AggregationEngine in aggregationEngine.go.

import (
	"sync"
	"time"
)

type MetricType int

const (
	Counter MetricType = iota
	Gauge
	Histogram
)

type Metric struct {
	Name      string
	Type      MetricType
	Value     float64
	Labels    map[string]string
	Timestamp time.Time
}

type MetricsCollector struct {
	mu       sync.RWMutex
	counters map[string]float64
	gauges   map[string]float64
	samples  map[string][]Sample
}

type Sample struct {
	Value     float64
	Timestamp time.Time
}

func NewMetricsCollector() *MetricsCollector {
	return &MetricsCollector{
		counters: make(map[string]float64),
		gauges:   make(map[string]float64),
		samples:  make(map[string][]Sample),
	}
}

func (mc *MetricsCollector) IncrementCounter(name string, value float64) {
	mc.mu.Lock()
	defer mc.mu.Unlock()
	mc.counters[name] += value
}

func (mc *MetricsCollector) SetGauge(name string, value float64) {
	mc.mu.Lock()
	defer mc.mu.Unlock()
	mc.gauges[name] = value
}

func (mc *MetricsCollector) RecordSample(name string, value float64) {
	mc.mu.Lock()
	defer mc.mu.Unlock()
	mc.samples[name] = append(mc.samples[name], Sample{
		Value:     value,
		Timestamp: time.Now(),
	})
}

func (mc *MetricsCollector) GetCounter(name string) float64 {
	mc.mu.RLock()
	defer mc.mu.RUnlock()
	return mc.counters[name]
}

func (mc *MetricsCollector) GetGauge(name string) float64 {
	mc.mu.RLock()
	defer mc.mu.RUnlock()
	return mc.gauges[name]
}

func (mc *MetricsCollector) GetSamples(name string) []Sample {
	mc.mu.RLock()
	defer mc.mu.RUnlock()
	result := make([]Sample, len(mc.samples[name]))
	copy(result, mc.samples[name])
	return result
}

func (mc *MetricsCollector) GetAllMetricNames() []string {
	mc.mu.RLock()
	defer mc.mu.RUnlock()
	names := make(map[string]bool)
	for k := range mc.counters {
		names[k] = true
	}
	for k := range mc.gauges {
		names[k] = true
	}
	for k := range mc.samples {
		names[k] = true
	}
	result := make([]string, 0, len(names))
	for k := range names {
		result = append(result, k)
	}
	return result
}
