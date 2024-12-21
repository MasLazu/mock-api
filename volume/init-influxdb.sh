# init-influxdb.sh
#!/bin/bash
# Wait for InfluxDB to be fully up
sleep 5

# Create the database
influx -execute "CREATE DATABASE k6"
