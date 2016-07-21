package main

import (
	"time"

	"github.com/convox/rack/api/provider"
	"github.com/convox/rack/api/workers"
)

func main() {
	go workers.StartAutoscale()
	go provider.MonitorCluster()
	go provider.MonitorHeartbeat()
	go provider.MonitorEvents()
	go workers.StartServicesCapacity() // This should be removed soon

	for {
		time.Sleep(1 * time.Hour)
	}
}
