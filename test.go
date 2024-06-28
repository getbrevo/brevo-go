package main

import (
	"context"
	"fmt"

	brevo "github.com/getbrevo/brevo-go/lib"
)

func main() {
	var ctx context.Context
	cfg := brevo.NewConfiguration()
	//Configure API key authorization: api-key
	cfg.AddDefaultHeader("api-key", "{api-key}")
	//Configure API key authorization: partner-key
	//cfg.AddDefaultHeader("partner-key", "YOUR_API_KEY")

	br := brevo.NewAPIClient(cfg)

	resp, err := br.ContactsApi.UpdateContact(ctx, "697", brevo.UpdateContact{
		Attributes: map[string]interface{}{
			"Email": "test-brevo@test.com",
		},
	})
	if err != nil {
		fmt.Println("Error  ", err.Error())
		return
	}
	fmt.Println("resp:", resp)
	//fmt.Printf("%+v", result.Contacts)
	//fmt.Println("=============\n\n\n", result.Contacts, result.Contacts.Failure, result.Contacts.Success)
	return
}
