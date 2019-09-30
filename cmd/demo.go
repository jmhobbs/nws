package main

import (
	"log"

	"github.com/jmhobbs/nws"
	"github.com/jmhobbs/nws/alerts"
	"github.com/jmhobbs/nws/alerts/status"
)

func main() {
	svc := alerts.New(nws.New("nws go library"))
	log.Println("==[ Alert Types ]==============")
	log.Println(svc.AlertTypes())

	log.Println("==[ Active Alerts ]============")
	active, err := svc.ActiveAlerts(alerts.ActiveAlertsOptions{
		AlertListOptions: alerts.AlertListOptions{
			Status: status.Actual,
			Area:   "NE",
		},
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, alert := range active.Graph {
		log.Println(alert.Headline)
	}

}
