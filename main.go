package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
)

const (
	SENDER   = "+14155238886"
	RECEIVER = "+4915901433811"
)

func getEnvVar(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}
	return os.Getenv(key)
}

func main() {
	twilioAccountSid := getEnvVar("TWILIO_ACCOUNT_SID")
	twilioAuthToken := getEnvVar("TWILIO_AUTH_TOKEN")
	_ = twilioAccountSid
	_ = twilioAuthToken

	fmt.Println("Welcome to the Household Planner!")

	client := twilio.NewRestClient()

	params := &api.CreateMessageParams{}
	params.SetFrom("whatsapp:" + SENDER)
	params.SetTo("whatsapp:" + RECEIVER)
	params.SetBody("Hier ist der Haushaltsplaner!\n")

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	} else {
		if resp.Body != nil {
			fmt.Println(*resp.Body)
		} else {
			fmt.Println(resp.Body)
		}
	}
}
