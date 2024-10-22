package main

import (
	"time"

	"github.com/marevers/power-datacenter-exporter/pkg/pdc"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promauto"

	log "github.com/sirupsen/logrus"
)

const (
	// LabelSerialNumber represents the inverter serial number
	LabelSerialNumber = "serialno"

	// Namespace is the metrics prefix
	Namespace = "pdc"
)

var (
	// Labels are the static labels that come with every metric
	labels = []string{
		LabelSerialNumber,
	}
)

type exporter struct {
	Reg     *prometheus.Registry
	Session *pdc.Session
	Metrics struct {
		GridFrequency1Vec *prometheus.GaugeVec
		GridFrequency2Vec *prometheus.GaugeVec
		GridVoltage1Vec   *prometheus.GaugeVec
		GridVoltage2Vec   *prometheus.GaugeVec

		PvInputVoltage1Vec *prometheus.GaugeVec
		PvInputVoltage2Vec *prometheus.GaugeVec
		PvInputCurrent1Vec *prometheus.GaugeVec
		PvInputCurrent2Vec *prometheus.GaugeVec

		AcOutputVoltage1Vec       *prometheus.GaugeVec
		AcOutputVoltage2Vec       *prometheus.GaugeVec
		AcOutputFrequency1Vec     *prometheus.GaugeVec
		AcOutputFrequency2Vec     *prometheus.GaugeVec
		AcOutputApparentPower1Vec *prometheus.GaugeVec
		AcOutputApparentPower2Vec *prometheus.GaugeVec
		AcOutputActivePower1Vec   *prometheus.GaugeVec
		AcOutputActivePower2Vec   *prometheus.GaugeVec

		OutputLoadPercent1Vec *prometheus.GaugeVec
		OutputLoadPercent2Vec *prometheus.GaugeVec

		BatVoltageVec       *prometheus.GaugeVec
		BatCapacityVec      *prometheus.GaugeVec
		BatChgCurrentVec    *prometheus.GaugeVec
		BatDischgCurrentVec *prometheus.GaugeVec

		TotalPvInputPowerVec          *prometheus.GaugeVec
		TotalOutputLoadPercentVec     *prometheus.GaugeVec
		TotalBatChgCurrentVec         *prometheus.GaugeVec
		TotalAcOutputApparentPowerVec *prometheus.GaugeVec
		TotalAcOutputActivePowerVec   *prometheus.GaugeVec

		ScrapeError prometheus.Gauge
	}
}

func createRegistry() *prometheus.Registry {
	reg := prometheus.NewRegistry()

	reg.MustRegister(
		collectors.NewGoCollector(),
		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
	)

	return reg
}

func (e *exporter) registerMetrics(labels []string) {
	// Grid

	e.Metrics.GridFrequency1Vec = promauto.With(e.Reg).NewGaugeVec(prometheus.GaugeOpts{
		Name:      "grid1_frequency",
		Namespace: Namespace,
		Help:      "Grid 1 frequency in herz",
	}, labels)

	e.Metrics.GridFrequency2Vec = promauto.With(e.Reg).NewGaugeVec(prometheus.GaugeOpts{
		Name:      "grid2_frequency",
		Namespace: Namespace,
		Help:      "Grid 2 frequency in herz",
	}, labels)

	e.Metrics.GridVoltage1Vec = promauto.With(e.Reg).NewGaugeVec(prometheus.GaugeOpts{
		Name:      "grid1_voltage",
		Namespace: Namespace,
		Help:      "Grid 1 voltage",
	}, labels)

	e.Metrics.GridVoltage2Vec = promauto.With(e.Reg).NewGaugeVec(prometheus.GaugeOpts{
		Name:      "grid2_voltage",
		Namespace: Namespace,
		Help:      "Grid 2 voltage",
	}, labels)

	// PV input

	e.Metrics.PvInputVoltage1Vec = promauto.With(e.Reg).NewGaugeVec(prometheus.GaugeOpts{
		Name:      "pvinput1_voltage",
		Namespace: Namespace,
		Help:      "PV input 1 voltage",
	}, labels)

	e.Metrics.PvInputVoltage2Vec = promauto.With(e.Reg).NewGaugeVec(prometheus.GaugeOpts{
		Name:      "pvinput2_voltage",
		Namespace: Namespace,
		Help:      "PV input 2 voltage",
	}, labels)

	e.Metrics.PvInputCurrent1Vec = promauto.With(e.Reg).NewGaugeVec(prometheus.GaugeOpts{
		Name:      "pvinput1_current",
		Namespace: Namespace,
		Help:      "PV input 1 current in amps",
	}, labels)

	e.Metrics.PvInputCurrent2Vec = promauto.With(e.Reg).NewGaugeVec(prometheus.GaugeOpts{
		Name:      "pvinput2_current",
		Namespace: Namespace,
		Help:      "PV input 2 current in amps",
	}, labels)

	// AC output

	e.Metrics.AcOutputVoltage1Vec = promauto.With(e.Reg).NewGaugeVec(prometheus.GaugeOpts{
		Name:      "acoutput1_voltage",
		Namespace: Namespace,
		Help:      "AC output 1 voltage",
	}, labels)

	e.Metrics.AcOutputVoltage2Vec = promauto.With(e.Reg).NewGaugeVec(prometheus.GaugeOpts{
		Name:      "acoutput2_voltage",
		Namespace: Namespace,
		Help:      "AC output 2 voltage",
	}, labels)

	e.Metrics.AcOutputFrequency1Vec = promauto.With(e.Reg).NewGaugeVec(prometheus.GaugeOpts{
		Name:      "acoutput1_frequency",
		Namespace: Namespace,
		Help:      "AC output 1 frequency in herz",
	}, labels)

	e.Metrics.AcOutputFrequency2Vec = promauto.With(e.Reg).NewGaugeVec(prometheus.GaugeOpts{
		Name:      "acoutput2_frequency",
		Namespace: Namespace,
		Help:      "AC output 2 frequency in herz",
	}, labels)

	e.Metrics.AcOutputApparentPower1Vec = promauto.With(e.Reg).NewGaugeVec(prometheus.GaugeOpts{
		Name:      "acoutput1_apparent_power",
		Namespace: Namespace,
		Help:      "AC output 1 apparent power in volt-amps",
	}, labels)

	e.Metrics.AcOutputApparentPower2Vec = promauto.With(e.Reg).NewGaugeVec(prometheus.GaugeOpts{
		Name:      "acoutput2_apparent_power",
		Namespace: Namespace,
		Help:      "AC output 2 apparent power in volt-amps",
	}, labels)

	e.Metrics.AcOutputActivePower1Vec = promauto.With(e.Reg).NewGaugeVec(prometheus.GaugeOpts{
		Name:      "acoutput1_active_power",
		Namespace: Namespace,
		Help:      "AC output 1 active power in watts",
	}, labels)

	e.Metrics.AcOutputActivePower2Vec = promauto.With(e.Reg).NewGaugeVec(prometheus.GaugeOpts{
		Name:      "acoutput2_active_power",
		Namespace: Namespace,
		Help:      "AC output 2 active power in watts",
	}, labels)

	// Output load

	e.Metrics.OutputLoadPercent1Vec = promauto.With(e.Reg).NewGaugeVec(prometheus.GaugeOpts{
		Name:      "output1_load_percent",
		Namespace: Namespace,
		Help:      "Output 1 load in percentage",
	}, labels)

	e.Metrics.OutputLoadPercent2Vec = promauto.With(e.Reg).NewGaugeVec(prometheus.GaugeOpts{
		Name:      "output2_load_percent",
		Namespace: Namespace,
		Help:      "Output 2 load in percentage",
	}, labels)

	// Battery

	e.Metrics.BatVoltageVec = promauto.With(e.Reg).NewGaugeVec(prometheus.GaugeOpts{
		Name:      "battery_voltage",
		Namespace: Namespace,
		Help:      "Battery voltage",
	}, labels)

	e.Metrics.BatCapacityVec = promauto.With(e.Reg).NewGaugeVec(prometheus.GaugeOpts{
		Name:      "battery_capacity_percent",
		Namespace: Namespace,
		Help:      "Battery capacity in percentage",
	}, labels)

	e.Metrics.BatChgCurrentVec = promauto.With(e.Reg).NewGaugeVec(prometheus.GaugeOpts{
		Name:      "battery_charge_current",
		Namespace: Namespace,
		Help:      "Battery charge current in amps",
	}, labels)

	e.Metrics.BatDischgCurrentVec = promauto.With(e.Reg).NewGaugeVec(prometheus.GaugeOpts{
		Name:      "battery_discharge_current",
		Namespace: Namespace,
		Help:      "Battery discharge current in amps",
	}, labels)

	// Totals

	e.Metrics.TotalPvInputPowerVec = promauto.With(e.Reg).NewGaugeVec(prometheus.GaugeOpts{
		Name:      "total_pvinput_power",
		Namespace: Namespace,
		Help:      "Total PV input power in watts",
	}, labels)

	e.Metrics.TotalOutputLoadPercentVec = promauto.With(e.Reg).NewGaugeVec(prometheus.GaugeOpts{
		Name:      "total_output_load_percent",
		Namespace: Namespace,
		Help:      "Total output load in percentage",
	}, labels)

	e.Metrics.TotalBatChgCurrentVec = promauto.With(e.Reg).NewGaugeVec(prometheus.GaugeOpts{
		Name:      "total_battery_charge_current",
		Namespace: Namespace,
		Help:      "Total battery charge current in amps",
	}, labels)

	e.Metrics.TotalAcOutputApparentPowerVec = promauto.With(e.Reg).NewGaugeVec(prometheus.GaugeOpts{
		Name:      "total_acoutput_apparent_power",
		Namespace: Namespace,
		Help:      "Total AC output apparent power in volt-amps",
	}, labels)

	e.Metrics.TotalAcOutputActivePowerVec = promauto.With(e.Reg).NewGaugeVec(prometheus.GaugeOpts{
		Name:      "total_acoutput_active_power",
		Namespace: Namespace,
		Help:      "Total AC output active power in watts",
	}, labels)

	// Scrape error

	e.Metrics.ScrapeError = promauto.With(e.Reg).NewGauge(prometheus.GaugeOpts{
		Name:      "scrape_error",
		Namespace: Namespace,
		Help:      "Returns 1 if the last scrape failed",
	})
}

func (e *exporter) calculateMetrics() error {
	e.Metrics.ScrapeError.Set(0)

	err := e.Session.GetWorkInfo()
	if err != nil {
		e.Metrics.ScrapeError.Set(1)
		return err
	}

	log.Infoln("Retrieved metrics from", e.Session.SerialNumber)

	var labelValues []string

	labelValues = append(
		labelValues,
		e.Session.SerialNumber,
	)

	e.Metrics.GridFrequency1Vec.WithLabelValues(labelValues...).Set(e.Session.WorkInfo.GridFrequency1)
	e.Metrics.GridFrequency2Vec.WithLabelValues(labelValues...).Set(e.Session.WorkInfo.GridFrequency2)
	e.Metrics.GridVoltage1Vec.WithLabelValues(labelValues...).Set(e.Session.WorkInfo.GridVoltage1)
	e.Metrics.GridVoltage2Vec.WithLabelValues(labelValues...).Set(e.Session.WorkInfo.GridVoltage2)

	e.Metrics.PvInputVoltage1Vec.WithLabelValues(labelValues...).Set(e.Session.WorkInfo.PvInputVoltage1)
	e.Metrics.PvInputVoltage2Vec.WithLabelValues(labelValues...).Set(e.Session.WorkInfo.PvInputVoltage2)
	e.Metrics.PvInputCurrent1Vec.WithLabelValues(labelValues...).Set(e.Session.WorkInfo.PvInputCurrent1)
	e.Metrics.PvInputCurrent2Vec.WithLabelValues(labelValues...).Set(e.Session.WorkInfo.PvInputCurrent2)

	e.Metrics.AcOutputVoltage1Vec.WithLabelValues(labelValues...).Set(e.Session.WorkInfo.AcOutputVoltage1)
	e.Metrics.AcOutputVoltage2Vec.WithLabelValues(labelValues...).Set(e.Session.WorkInfo.AcOutputVoltage2)
	e.Metrics.AcOutputFrequency1Vec.WithLabelValues(labelValues...).Set(e.Session.WorkInfo.AcOutputFrequency1)
	e.Metrics.AcOutputFrequency2Vec.WithLabelValues(labelValues...).Set(e.Session.WorkInfo.AcOutputFrequency2)
	e.Metrics.AcOutputApparentPower1Vec.WithLabelValues(labelValues...).Set(e.Session.WorkInfo.AcOutputApparentPower1)
	e.Metrics.AcOutputApparentPower2Vec.WithLabelValues(labelValues...).Set(e.Session.WorkInfo.AcOutputApparentPower2)
	e.Metrics.AcOutputActivePower1Vec.WithLabelValues(labelValues...).Set(e.Session.WorkInfo.AcOutputActivePower1)
	e.Metrics.AcOutputActivePower2Vec.WithLabelValues(labelValues...).Set(e.Session.WorkInfo.AcOutputActivePower2)

	e.Metrics.OutputLoadPercent1Vec.WithLabelValues(labelValues...).Set(e.Session.WorkInfo.OutputLoadPercent1)
	e.Metrics.OutputLoadPercent2Vec.WithLabelValues(labelValues...).Set(e.Session.WorkInfo.OutputLoadPercent2)

	e.Metrics.BatVoltageVec.WithLabelValues(labelValues...).Set(e.Session.WorkInfo.BatVoltage)
	e.Metrics.BatCapacityVec.WithLabelValues(labelValues...).Set(e.Session.WorkInfo.BatCapacity)
	e.Metrics.BatChgCurrentVec.WithLabelValues(labelValues...).Set(e.Session.WorkInfo.BatChgCurrent)
	e.Metrics.BatDischgCurrentVec.WithLabelValues(labelValues...).Set(e.Session.WorkInfo.BatDischgCurrent)

	e.Metrics.TotalPvInputPowerVec.WithLabelValues(labelValues...).Set(e.Session.WorkInfo.TotalPvInputPower)
	e.Metrics.TotalOutputLoadPercentVec.WithLabelValues(labelValues...).Set(e.Session.WorkInfo.TotalOutputLoadPercent)
	e.Metrics.TotalBatChgCurrentVec.WithLabelValues(labelValues...).Set(e.Session.WorkInfo.TotalBatChgCurrent)
	e.Metrics.TotalAcOutputApparentPowerVec.WithLabelValues(labelValues...).Set(e.Session.WorkInfo.TotalAcOutputApparentPower)
	e.Metrics.TotalAcOutputActivePowerVec.WithLabelValues(labelValues...).Set(e.Session.WorkInfo.TotalAcOutputActivePower)

	return nil
}

func startMetricsTicker(e *exporter, t time.Duration) {
	tck := time.NewTicker(t)
	defer tck.Stop()

	for {
		select {
		case <-tck.C:
			err := e.calculateMetrics()
			if err != nil {
				log.Warnln(err)
			}
		}
	}
}
