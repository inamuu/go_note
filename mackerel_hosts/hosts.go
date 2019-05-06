package main

import (
	"fmt"
	"os"

	mkr "github.com/mackerelio/mackerel-client-go"
)

func main() {
	client := mkr.NewClient(os.Getenv("MACKEREL_APIKEY"))
	hosts, _ := client.FindHosts(&mkr.FindHostsParam{
		Statuses: []string{"working"},
	})

	for _, v := range hosts {
		fmt.Printf(" HOSTID: %v\n Name: %v\n\n", v.ID, v.Name)
	}
}