import http.server
import socketserver
from prometheus_client import start_http_server, Summary

APP_PORT = 5000
PROMETHEUS_PORT = 8080

start_http_server(PROMETHEUS_PORT)
print("Prometheus exporter server started at localhost:" + str(PROMETHEUS_PORT))

REQUEST_TIME = Summary('http_request_time_spent', 'Time spent processing HTTP request')
class InstrumentedHttpRequestHandler(http.server.SimpleHTTPRequestHandler):
    @REQUEST_TIME.time()
    def do_HEAD(self):
        return http.server.SimpleHTTPRequestHandler.do_HEAD(self)
    @REQUEST_TIME.time()
    def do_GET(self):
        return http.server.SimpleHTTPRequestHandler.do_GET(self)
    @REQUEST_TIME.time()
    def do_POST(self):
        return http.server.SimpleHTTPRequestHandler.do_POST(self)

with socketserver.TCPServer(("", APP_PORT), InstrumentedHttpRequestHandler) as httpd:
    print("Static content HTTP server started at localhost:" + str(APP_PORT))
    httpd.serve_forever()
