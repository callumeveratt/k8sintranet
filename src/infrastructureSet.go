package main

import (
	"html/template"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

// Global so can be used cross function
var grafanaURL string
var prometheusURL string
var jaegerURL string
var linkerdURL string
var rabbitMQURL string
var alertManagerURL string
var linkerdEnabled string
var rabbitMQEnabled string
var alertManagerEnabled string
var grafanaEnabled string
var jaegerEnabled string
var prometheusEnabled string

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("site/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Error(err.Error())
	}

	data := struct {
		Grafana             string
		GrafanaEnabled      string
		Linkerd             string
		LinkerdEnabled      string
		RabbitMQ            string
		RabbitMQEnabled     string
		AlertManager        string
		AlertManagerEnabled string
		Jaeger              string
		JaegerEnabled       string
		Prometheus          string
		PrometheusEnabled   string
	}{
		Grafana:             grafanaURL,
		GrafanaEnabled:      grafanaEnabled,
		Linkerd:             linkerdURL,
		LinkerdEnabled:      linkerdEnabled,
		RabbitMQ:            rabbitMQURL,
		RabbitMQEnabled:     rabbitMQEnabled,
		AlertManager:        alertManagerURL,
		AlertManagerEnabled: alertManagerEnabled,
		Jaeger:              jaegerURL,
		JaegerEnabled:       jaegerEnabled,
		Prometheus:          prometheusURL,
		PrometheusEnabled:   prometheusEnabled,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Error(err.Error())
	}
}

func main() {
	// Set Logging
	log.SetFormatter(&log.JSONFormatter{})
	var logLevel = os.Getenv("logLevel")
	log.Info("Log level is " + logLevel)

	if logLevel == "debug" {
		log.SetLevel(log.DebugLevel)
	}

	// Get links from environment
	grafanaEnabled = os.Getenv("grafanaEnabled")
	jaegerEnabled = os.Getenv("jaegerEnabled")
	prometheusEnabled = os.Getenv("prometheusEnabled")
	linkerdEnabled = os.Getenv("linkerdEnabled")
	rabbitMQEnabled = os.Getenv("rabbitMQEnabled")
	alertManagerEnabled = os.Getenv("alertManagerEnabled")

	if grafanaEnabled == "true" {
		grafanaURL = os.Getenv("grafanaURL")
		log.Info("Grafana is enabled, URL " + grafanaURL)
	}

	if prometheusEnabled == "true" {
		prometheusURL = os.Getenv("prometheusURL")
		log.Info("Prometheus is enabled, URL " + prometheusURL)
	}

	if jaegerEnabled == "true" {
		jaegerURL = os.Getenv("jaegerURL")
		log.Info("Jaeger is enabled, URL " + jaegerURL)
	}

	if linkerdEnabled == "true" {
		linkerdURL = os.Getenv("linkerdURL")
		log.Info("Linkerd is enabled, URL " + linkerdURL)
	}

	if rabbitMQEnabled == "true" {
		rabbitMQURL = os.Getenv("rabbitMQURL")
		log.Info("RabbitMQ is enabled, URL " + rabbitMQURL)
	}

	if alertManagerEnabled == "true" {
		alertManagerURL = os.Getenv("alertManagerURL")
		log.Info("AlertManager is enabled, URL " + alertManagerURL)
	}

	var port = os.Getenv("port")
	log.Info("Will be listening on port " + port)
	log.Info("Starting HTTP Server")
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	log.Info(http.ListenAndServe(":"+port, mux))
}
