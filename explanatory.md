# Beginner Explanatory Guide: PLATFORM-2918: Build custom metrics aggregation engine

> **Task Type**: Product Task  
> **Domain/Focus**: Backend Development, Metrics Aggregation, Golang

---

## 1. The Goal (In-Depth Beginner Explanation)

### The Core Problem
In the current state of our application, we have a `MetricsCollector` that can register and increment counters, which means it can keep track of how many times certain events occur. However, it lacks the ability to aggregate these metrics into meaningful insights. This is a significant limitation because without aggregation, we cannot analyze trends, understand performance, or make informed decisions based on the data collected. For instance, if we only know that a certain event happened 100 times, we cannot determine how often it happens over time or how it compares to other events.

The task at hand is to implement an `AggregationEngine` that will allow us to compute various statistical metrics such as percentiles (e.g., p50, p90, p99), rates (how many events occur per second), and histograms (which categorize data into ranges). This functionality is crucial for observability, as it enables us to monitor system performance and user behavior effectively. By fixing this issue, we will enhance our application's ability to provide insights, leading to better performance tuning and user experience.

### Jargon Buster (Key Terms Explained)
* **Aggregation**: This is the process of combining multiple data points into a single summary statistic. For example, if we have a list of response times for a web service, aggregation can help us find the average response time or the maximum response time. Without aggregation, we would only have raw data without context.

* **Percentiles**: Percentiles are measures that indicate the value below which a given percentage of observations fall. For instance, the 90th percentile (p90) means that 90% of the data points are below this value. This is useful for understanding the distribution of data, especially in performance metrics.

* **Rate**: In the context of metrics, a rate typically refers to the frequency of events occurring over a specific time period. For example, if a service processes 300 requests in 10 seconds, the rate would be 30 requests per second. This helps in understanding how busy a service is over time.

* **Histogram**: A histogram is a graphical representation that organizes a group of data points into user-specified ranges (or bins). For example, if we want to analyze response times, we might create bins for 0-100ms, 100-200ms, etc. This allows us to see how many requests fall into each time range.

### Expected Outcome
After implementing the `AggregationEngine`, the system should be able to perform the following functions:

- **Before**: The `MetricsCollector` can only count occurrences of events without any statistical analysis.
- **After**: The `AggregationEngine` will allow the application to:
  - Record metric values with timestamps using `AddSample()`.
  - Compute percentiles (p50, p90, p99) using `GetPercentile()`.
  - Calculate the rate of events over a time window with `GetRate()`.
  - Create histograms of samples with `GetHistogram()`.
  - Reset samples after exporting them using `Flush()`.

This enhancement will provide a richer set of tools for analyzing metrics, leading to better insights and decision-making.

---

## 2. Related Coding Concepts & Syntax (50% Theory, 50% Practice)

### Concept 1: Structs in Golang
#### 📘 Theoretical Overview (50%)
* **Why it exists**: In Golang, structs are used to create complex data types that group together variables (fields) under a single name. This is essential for organizing data in a way that reflects real-world entities. Without structs, managing related data would be cumbersome and error-prone.

* **Key Mechanisms**: A struct is defined using the `type` keyword followed by the struct name and its fields. Each field can have a different data type. Structs can also have methods associated with them, allowing for encapsulation of behavior along with data.

#### 💻 Syntax & Practical Examples (50%)
* **Language Syntax**:
  ```go
  type MetricSample struct {
      Value     float64
      Timestamp time.Time
  }
  ```
  - `type MetricSample`: This defines a new struct type named `MetricSample`.
  - `Value float64`: This field stores the metric value as a floating-point number.
  - `Timestamp time.Time`: This field stores the time when the metric was recorded.

* **Real-World Application**:
  ```go
  func NewMetricSample(value float64) MetricSample {
      return MetricSample{
          Value:     value,
          Timestamp: time.Now(),
      }
  }
  ```
  - This function creates a new `MetricSample` with the current time as the timestamp.

---

## 3. Step-by-Step Logic & Walkthrough

1. **Step 1: Locate and Analyze the Target File**
   * Navigate to the `p-w05-task-04` folder and open the file named `AggregationEngine.go`.
   * Look for the `AddSample()` function, which is where we will implement the logic to record metric values.

2. **Step 2: Input Verification & Validation**
   * Before adding a sample, check if the input value is valid (e.g., not negative or NaN). This ensures that only meaningful data is recorded.

3. **Step 3: Core Implementation / Modification**
   * Implement the `AddSample()` function to append the new metric sample to an internal slice of samples. Ensure that each sample includes a timestamp.

4. **Step 4: Output Verification & Testing**
   * After implementing the function, run the unit tests provided in the repository to ensure that all tests pass. This will confirm that the logic is correct and that the new functionality works as expected.

---

## 4. Detailed Walkthrough of Test Cases

### Test Case 1: Standard / Success Case
* **Description**: This test checks if the `AddSample()` function correctly records a valid metric sample.
* **Inputs**:
  ```json
  {
      "value": 42.0
  }
  ```
* **Step-by-Step Execution Trace**:
  1. The input value `42.0` is received by the `AddSample()` function.
  2. The function checks if the value is valid (not negative or NaN).
  3. The function appends a new `MetricSample` to the internal slice with the current timestamp.
  4. The function completes without errors.
* **Expected Output**: The internal slice of samples now contains one entry with the value `42.0` and the corresponding timestamp.

### Test Case 2: Edge Case / Validation Fail
* **Description**: This test checks how the system handles an invalid input (negative value).
* **Inputs**:
  ```json
  {
      "value": -5.0
  }
  ```
* **Step-by-Step Execution Trace**:
  1. The input value `-5.0` is received by the `AddSample()` function.
  2. The validation block detects that the input is invalid (negative).
  3. The execution is halted early, and an error is returned.
  4. No new sample is added to the internal slice.
* **Expected Output**: The function returns an error indicating that the input value cannot be negative, and the internal slice remains unchanged.