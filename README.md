# Golang HTTP server with Prometheus instrumentation

This image can be used to serve static content over HTTP, while exporting metrics in Prometheus format.

HTTP requests are instrumented with metrics to count the total amount of time spent serving HTTP requests, as well as the number of requests received. These metrics are all exposed on a separate "metrics" port for collection, in Prometheus format.

This image will serve static content from the `/www` directory on port `5000` by default.
Prometheus metrics are exposed on port `8080` by default.

### How to use

This docker container is hosted using Github's package registry. Images can be pulled as follows:

`docker pull docker.pkg.github.com/joeyx22lm/static-http-server-prometheus/static-http-server-prometheus:v0.2.0`

### Environment Variables
|      Variable      |                    Description                 | Default |
|--------------------|------------------------------------------------|---------|
|   PROMETHEUS_PORT  | Port to expose Prometheus metrics server       |  8080   |
|         PORT       | Port to expose HTTP server with static content |  5000   |


### Changelog

`v1.0.0`: A Golang-based HTTP server that exposes Prometheus metrics.
 - Base image: `golang:1.16.3-alpine3.13`
 - Tag(s): `v1.0.0`, `latest`

`v0.1.0`: A Python-based HTTP server that exposes Prometheus metrics.
 - Base image: `python:3.9.4-buster`
 - Tag(s): `v0.1.0`
