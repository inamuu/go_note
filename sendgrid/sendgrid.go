package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sendgrid/sendgrid-go"
)

func main() {
	request := sendgrid.GetRequest(os.Getenv("SENDGRID_API_KEY"), "/v3/api_keys", "https://api.sendgrid.com")
	request.Method = "GET"

	response, err := sendgrid.API(request)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}
