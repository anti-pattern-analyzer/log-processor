# Log Processor Overview

This document provides an overview of the **Log Processor**, a Go-based service that continuously consumes, processes, and structures logs from microservices. It is designed to work closely with Kafka (as a log stream), TimeSeries/relational databases (for structured data storage), and optionally Neo4j (for visualizing microservice dependencies).

---

## Table of Contents

1. [Purpose](#purpose)  
2. [High-Level Architecture](#high-level-architecture)  
3. [Key Components](#key-components)  
4. [Data Flow](#data-flow)  
5. [Typical Use Cases](#typical-use-cases)  
6. [Pitfalls and Recommendations](#pitfalls-and-recommendations)  
7. [Extensibility](#extensibility)  
8. [Contributing](#contributing)  

---

## Purpose

The Log Processor’s primary goals:

- **Ingest and store** raw logs from Kafka topics (published by microservices).
- **Reconstruct** complete request flows (traces) by correlating data with `trace_id` and `span_id`.
- **Convert** raw data into a structured format, enabling easy querying, analytics, and continuous updates to runtime models.
- **Periodically update** a service dependency graph (if desired, in a graph database like Neo4j).

---

## High-Level Architecture

1. **Kafka Consumer**  
   - Reads log messages published by microservices to Kafka.  
   - Each message typically includes:
     - `trace_id` / `span_id`
     - Source service, destination service
     - Request/response details
     - Timestamps
     - Possible parent–child relationships for hierarchical call chains

2. **Database Storage**  
   - A time-series or relational database is used for:
     - **Raw Logs**: storing individual log records as they arrive.
     - **Structured Logs**: storing aggregated records (grouped by `trace_id` to rebuild full request flows).

3. **Dependency Graph** (Optional)  
   - A scheduler periodically pulls processed data from the database and updates a dependency graph, capturing:
     - Service relationships (which service calls which)
     - Call metrics (call frequency, average latency, error counts)

4. **Dashboard / Analytics**  
   - Analytical dashboards (e.g., Grafana) or custom UIs can query the Log Processor’s outputs for real-time observability.

---

## Key Components

![image](https://github.com/user-attachments/assets/c06cf475-baa4-4118-bd8c-2ee4700cf8a7)

### 1. Kafka Consumer
- **Language**: Go  
- **Responsibility**:  
  - Subscribes to one or more Kafka topics (e.g. `microservice-logs-topic`).
  - Batches or streams log entries to the Log Processor’s core logic.

### 2. Processing / Enrichment
- **Description**:  
  - Groups log entries by `trace_id`.
  - Reconstructs call hierarchies using parent–child `span_id`.
  - Performs transformations or adds metadata (e.g. user sessions, environment markers).

### 3. Time-Series / Relational Database
- **Purpose**:  
  - Holds both raw and enriched log data for querying.
  - Typically includes tables for `raw_logs` and `structured_logs`, keyed by `trace_id`.

### 4. Dependency Graph Updater
- **Purpose**:  
  - Reads from `structured_logs` in small batches (e.g., every few seconds/minutes).
  - Inserts or updates edges in a graph database to reflect new inter-service calls or changes in latencies.

---

## Data Flow

1. **Microservice -> Kafka**  
   Each microservice logs events to Kafka, tagging each entry with a `trace_id` and `span_id` for correlation.

2. **Kafka -> Raw Logs Table**  
   The Log Processor consumes messages from Kafka and writes them to a `raw_logs` table. This ensures minimal overhead during ingestion.

3. **Log Processor -> Structured Logs**  
   After correlation and trace reconstruction, the processor writes aggregated data (per trace) to `structured_logs`. This can include total duration, involved services, and error info.

4. **Structured Logs -> Dependency Graph**  
   The processor (or a background job) uses the aggregated data to update a dependency graph, capturing service-to-service calls and associated metrics.

---

## Typical Use Cases

- **Real-Time Monitoring**  
  Continuously track call frequency, average/95th percentile latency, error rates, and feed data into dashboards or alerts.

- **Distributed Tracing**  
  Reconstruct full call flows from `trace_id` and `span_id`, enabling quick identification of bottlenecks or anomalies.

- **Architecture Insights**  
  Maintain an always-fresh overview of how services interact at runtime. Detect cyclical dependencies, newly introduced services, or changes in network topologies.

---

## Pitfalls and Recommendations

1. **High-Volume Logging**  
   - With many services, logs can be massive. Consider partitioning Kafka topics and scaling the Log Processor horizontally.

2. **Indexing**  
   - Ensure critical columns (e.g. `trace_id`, `span_id`) are properly indexed in your database to prevent query bottlenecks.

3. **Partial Traces**  
   - Logs may arrive out-of-order or may be missing. Handle incomplete data gracefully (e.g., partial correlation, skipping incomplete spans).

4. **Fault Tolerance**  
   - Use replication for Kafka or database.  
   - Log Processor itself should be stateless, allowing easy re-deployment or autoscaling.

---

## Extensibility

- **Custom Enrichment**  
  Add domain-specific metadata (e.g., attach a user token, geo-location) to each structured log entry.

- **Alternative Storage**  
  Replace or add other data stores (e.g. Elasticsearch, InfluxDB) for specialized indexing or analytics demands.

- **Integration with Analytics**  
  Pipe your structured logs into ML pipelines or anomaly detection services for advanced performance insights.

---

## Contributing

1. **Fork the repository** and create a feature branch.
2. **Implement and test** changes or bug fixes.
3. **Open a Pull Request** describing your changes.

---
