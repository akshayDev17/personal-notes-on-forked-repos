package main

import (
	"context"
	"fmt"
	"os"

	"github.com/plutov/paypal/v4"
)

func main() {

	// Create a client instance
	c, err := paypal.NewClient("clientID", "secretID", paypal.APIBaseSandBox)
	c.SetLog(os.Stdout) // Set log to terminal stdout

	_, err = c.GetAccessToken(context.Background())
	if err != nil {
		fmt.Println(err)
	}
}
