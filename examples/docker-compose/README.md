# Docker Compose Example

This is just an example setup for the exporter in a small stack with Grafana and Prometheus.

Create `.env` file with the following content:

```
PDC_USERNAME=<your power-datacenter username>
PDC_PASSWORD=<your power-datacenter password>
PDC_SERIALNUMBER=<serial number of the device to monitor>
GRAFANA_ADMIN_PASSWORD=<admin password for grafana>
```

Start compose setup:

```sh
docker compose up -d
```
