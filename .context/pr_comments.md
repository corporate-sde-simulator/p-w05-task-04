# PR Review - Custom metrics collector and aggregator (by Deepak)

## Reviewer: Priya Menon
---

**Overall:** Good foundation but critical bugs need fixing before merge.

### `metricsCollector.go`

> **Bug #1:** Histogram bucket boundaries are inclusive on both sides causing double-counting at boundaries
> This is the higher priority fix. Check the logic carefully and compare against the design doc.

### `aggregationEngine.go`

> **Bug #2:** Timer metric records duration in nanoseconds but reports as milliseconds without conversion
> This is more subtle but will cause issues in production. Make sure to add a test case for this.

---

**Deepak**
> Acknowledged. I have documented the issues for whoever picks this up.
