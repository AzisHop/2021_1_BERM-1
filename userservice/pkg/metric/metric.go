package metric

import (
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
	"strconv"
)

const ctxKeyStartReqTime uint8 = 5

var (
	hits    *prometheus.CounterVec
	errors  *prometheus.CounterVec
	Timings *prometheus.SummaryVec
)

func Destroy() {
	prometheus.Unregister(hits)
	prometheus.Unregister(errors)
	prometheus.Unregister(Timings)
}

func New() {
	hits = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "hits",
	}, []string{"status", "path"})
	errors = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "errors",
	}, []string{"error"})

	Timings = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name:       "timings",
			Help:       "Timer running action",
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		},
		[]string{"method", "URL"})
	prometheus.MustRegister(hits, errors, Timings)
}


func CrateRequestHits(status int, r *http.Request) {
	route := mux.CurrentRoute(r)
	path, _ := route. GetPathTemplate()
	hits.WithLabelValues(strconv.Itoa(status), path).Inc()
}

func CrateRequestError(err error) {
	if err != nil {
		errors.WithLabelValues(err.Error()).Inc()
	}
}
