package main

import (
    "fmt"
    "log"
    "io"
    "net/http"
    "os"
    "github.com/prometheus/client_golang/prometheus/promhttp"
    metrics "github.com/slok/go-http-metrics/metrics/prometheus"
    "github.com/slok/go-http-metrics/middleware"
    middlewarestd "github.com/slok/go-http-metrics/middleware/std"
)

type fileOnlyFilesystem struct {
    fs               http.FileSystem
    readDirBatchSize int
}

func (fs fileOnlyFilesystem) Open(name string) (http.File, error) {
    f, err := fs.fs.Open(name)
    if err != nil {
        return nil, err
    }
    return neuteredStatFile{File: f, readDirBatchSize: fs.readDirBatchSize}, nil
}

type neuteredStatFile struct {
    http.File
    readDirBatchSize int
}

func (e neuteredStatFile) Stat() (os.FileInfo, error) {
    s, err := e.File.Stat()
    if err != nil {
        return nil, err
    }
    if s.IsDir() {
    LOOP:
        for {
            fl, err := e.File.Readdir(e.readDirBatchSize)
            switch err {
            case io.EOF:
                break LOOP
            case nil:
                for _, f := range fl {
                    if f.Name() == "index.html" {
                        return s, err
                    }
                    if f.Name() == "index.htm" {
                        return s, err
                    }
                }
            default:
                return nil, err
            }
        }
        return nil, os.ErrNotExist
    }
    return s, err
}

func main() {
    prometheus_port := ":" + os.Getenv("PROMETHEUS_PORT")
    http_port := ":" + os.Getenv("PORT")

    go func() {
        fmt.Println("Prometheus server listening on ", prometheus_port)
        log.Fatal(http.ListenAndServe(prometheus_port, promhttp.Handler()))
    }()

    // HTTP metrics instrumentation middleware.
    mdlw := middleware.New(middleware.Config{
        Recorder: metrics.NewRecorder(metrics.Config{}),
    })

    fmt.Println("FileServer listening on ", http_port)
    fs := fileOnlyFilesystem{fs: http.Dir("/www/"), readDirBatchSize: 2}
    h := middlewarestd.Handler("", mdlw, http.FileServer(fs))
    http.ListenAndServe(http_port, h)
}
