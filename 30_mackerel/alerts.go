package main

import (
	"fmt"
	"os"

	mkr "github.com/mackerelio/mackerel-client-go"
)

func main() {
	client := mkr.NewClient(os.Getenv("MACKEREL_APIKEY"))
	alerts, err := client.FindAlerts()
	if err != nil {
		fmt.Println("no alerts")
		os.Exit(0)
	}

	for _, res := range alerts.Alerts {
		fmt.Printf(" AlertType: %v\n AlertHostID: %v\n AlertStatus: %v\n", res.Type, res.HostID, res.Status)
	}
}
