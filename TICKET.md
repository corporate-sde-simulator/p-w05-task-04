# PLATFORM-2918: Build custom metrics aggregation engine

**Status:** In Progress · **Priority:** High
**Sprint:** Sprint 27 · **Story Points:** 5
**Reporter:** Vikram Patel (Observability Lead) · **Assignee:** You (Intern)
**Due:** End of sprint (Friday)
**Labels:** `backend`, `golang`, `observability`, `metrics`
**Task Type:** Feature Ship

---

## Description

The metrics collector can register and increment counters, but has **no aggregation capability**. We need an aggregation engine that computes percentiles, rates, and histograms from raw metric samples. The working `MetricsCollector` exists. Implement the TODO items in `AggregationEngine`.

## Acceptance Criteria

- [ ] `AddSample()` records a metric value with timestamp
- [ ] `GetPercentile()` computes p50, p90, p99 from recorded samples
- [ ] `GetRate()` computes per-second rate over a time window
- [ ] `GetHistogram()` buckets samples into configurable ranges
- [ ] `Flush()` resets samples after export
- [ ] All unit tests pass
