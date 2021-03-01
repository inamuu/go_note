package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {

	url := "https://api.sendgrid.com/v3/marketing/contacts"
	apiKey := os.Getenv("SENDGRID_API_KEY")
	list_id := os.Getenv("LIST_ID")
	email := os.Getenv("EMAIL")

	if apiKey == "" {
		fmt.Println("Set API Key")
		return
	}

	if list_id == "" {
		fmt.Println("Set List Id")
		return
	}

	payload := strings.NewReader(fmt.Sprintf(`{
		"list_ids":["%s"],
		"contacts":[{"email":"%s"}]
		}`, list_id, email))

	req, _ := http.NewRequest("PUT", url, payload)

	req.Header.Add("authorization", "Bearer "+fmt.Sprintf(apiKey))
	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
