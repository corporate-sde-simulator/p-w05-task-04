package metrics

// AggregationEngine — computes statistical aggregates from raw metric samples.
//
// YOU MUST IMPLEMENT the methods marked with TODO below.
// The MetricsCollector in metricsCollector.go is working — it feeds samples here.

import (
	"sort"
	"time"
)

type AggregationEngine struct {
	collector *MetricsCollector
}

type HistogramBucket struct {
	LowerBound float64
	UpperBound float64
	Count      int
}

type AggregationResult struct {
	MetricName string
	SampleCount int
	Min         float64
	Max         float64
	Sum         float64
	Average     float64
}

func NewAggregationEngine(collector *MetricsCollector) *AggregationEngine {
	return &AggregationEngine{collector: collector}
}

// GetPercentile computes the p-th percentile (0-100) for a named metric.
//
// TODO: Implement this method.
// 1. Get samples from collector using GetSamples(name)
// 2. Extract values and sort them ascending
// 3. Calculate the index: idx = (percentile / 100) * (len - 1)
// 4. If idx is whole number, return that value
// 5. If fractional, interpolate between floor(idx) and ceil(idx)
// 6. Return 0 if no samples exist
func (ae *AggregationEngine) GetPercentile(name string, percentile float64) float64 {
	// TODO: Implement percentile calculation
	return 0
}

// GetRate computes the per-second rate of a metric over the given duration.
//
// TODO: Implement this method.
// 1. Get samples from collector
// 2. Filter to only samples within the time window (now - duration, now)
// 3. Count the number of samples in the window
// 4. Divide count by duration in seconds
// 5. Return 0 if no samples or duration is 0
func (ae *AggregationEngine) GetRate(name string, duration time.Duration) float64 {
	// TODO: Implement rate calculation
	return 0
}

// GetHistogram buckets samples into ranges defined by boundaries.
// boundaries example: [0, 10, 50, 100, 500] creates buckets [0-10), [10-50), [50-100), [100-500), [500+)
//
// TODO: Implement this method.
// 1. Get samples from collector
// 2. Sort boundaries ascending
// 3. Create buckets from consecutive boundary pairs
// 4. For each sample value, find which bucket it belongs to
// 5. Return slice of HistogramBucket with counts
func (ae *AggregationEngine) GetHistogram(name string, boundaries []float64) []HistogramBucket {
	// TODO: Implement histogram bucketing
	return nil
}

// GetSummary returns basic statistics (min, max, sum, avg, count) for a metric.
//
// TODO: Implement this method.
// 1. Get samples from collector
// 2. Compute min, max, sum from values
// 3. Compute average = sum / count
// 4. Return AggregationResult struct
func (ae *AggregationEngine) GetSummary(name string) AggregationResult {
	// TODO: Implement summary statistics
	return AggregationResult{MetricName: name}
}

// Flush clears all samples from the collector after aggregation.
func (ae *AggregationEngine) Flush(name string) {
	// This is provided — no TODO needed
	_ = sort.Float64s // imported for student use
}
