package main

import (
	"flag"
	"net/http"
	"time"

	"github.com/marevers/power-datacenter-exporter/pkg/pdc"

	log "github.com/sirupsen/logrus"
)

var (
	logLevel     = flag.String("log.level", "info", "Log level for logging.")
	listenAddr   = flag.String("web.listen-address", ":8080", "The address to listen on for HTTP requests.")
	metricsPath  = flag.String("web.telemetry-path", "/metrics", "Path under which to expose metrics.")
	baseUrl      = flag.String("pdc.baseurl", "", "Base URL to use.")
	username     = flag.String("pdc.username", "", "Username for logging in.")
	password     = flag.String("pdc.password", "", "Password for logging in.")
	serialNumber = flag.String("pdc.serialnumber", "", "Serial number of device.")
	interval     = flag.Int("pdc.interval", 60, "Interval in seconds for data polling.")
)

func main() {
	flag.Parse()

	if level, err := log.ParseLevel(*logLevel); err != nil {
		log.Fatalln(err)
	} else {
		log.SetLevel(level)
	}

	ses := pdc.NewSession(*baseUrl, *serialNumber)

	if err := ses.Login(*username, *password); err != nil {
		log.Fatalln(err)
	}

	exporter := &exporter{
		Reg:     createRegistry(),
		Session: ses,
	}

	exporter.registerMetrics(labels)

	srv := &http.Server{
		Addr:    *listenAddr,
		Handler: exporter.routes(),
	}

	go startMetricsTicker(exporter, time.Duration(*interval)*time.Second)

	log.Println("Starting power-datacenter Exporter at", *listenAddr)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalln("Error starting HTTP server:", err)
	}
}
