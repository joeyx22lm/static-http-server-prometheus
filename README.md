# Golang HTTP server with Prometheus instrumentation

This image can be used to serve static content over HTTP, while exporting metrics in Prometheus format.

All `GET`, `HEAD`, and `POST` requests are instrumented with a `Summary` metric to count the total amount of time spent serving HTTP requests, as well as the number of requests received. These metrics are all exposed on a separate "metrics" port for collection, in Prometheus format.

This image will serve static content from the `/www` directory on port `5000` by default.
Prometheus metrics are exposed on port `8080` by default.

### Environment Variables

|      Variable      |                 Description               | Default |
|--------------------|-------------------------------------------|---------|
|   PROMETHEUS_PORT  | Port to expose Prometheus metrics server  |  8080   |
|         PORT       | Port to expose HTTP server                |  5000   |
