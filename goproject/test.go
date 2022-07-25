package main

import (
	"context"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func main() {
	bucket := "test"
	org := "ISCAS"
	token := "V_6KyJzzVWMMErhXzUExm2GgU9mQ5DZYfVNgf6wTLv8BknwVbCCM21lY9k2209PEe7E0Gt_Um2HUD4aKvIn-6w=="
	url := "http://localhost:8086"
	client := influxdb2.NewClient(url, token)
	writeAPI := client.WriteAPIBlocking(org, bucket)
	p := influxdb2.NewPoint("stat",
		map[string]string{"unit": "temperature"},
		map[string]interface{}{"avg": 24.5, "max": 45},
		time.Now())
	writeAPI.WritePoint(context.Background(), p)
	client.Close()
}
