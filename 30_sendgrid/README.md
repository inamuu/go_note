Test Sendgrid API v3 using sendgrid-go
---

I want to add a contact to my contact list of sendgrid marketing campaign.

ref: https://github.com/sendgrid/sendgrid-go#processing-inbound-email

To add a contact using 

ref: https://sendgrid.com/docs/api-reference/

## How to use

1. Get API key from SendGrid Web Console.
2. Run below command

```sh
SENDGRID_API_KEY=SG.XXXXX LIST_ID=XXX EMAIL=XXX go run ./sendgrid.go
```
