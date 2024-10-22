# power-datacenter Exporter

This exporter allows you to retrieve solar inverter / battery statistics from inverters connected to `power-datacenter.com` and convert them into Prometheus metrics for use within Prometheus rules or Grafana Dashboards.

Note: the statistics are only updated once every 5 minutes, so scraping more often than that does not result in higher resolution metrics.
