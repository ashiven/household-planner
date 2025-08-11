package planner

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
)

var SENDER = getEnvVar("WHATSAPP_SENDER")

func getEnvVar(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}
	return os.Getenv(key)
}

func InitializeTwilioClient() *twilio.RestClient {
	twilioAccountSid := getEnvVar("TWILIO_ACCOUNT_SID")
	twilioAuthToken := getEnvVar("TWILIO_AUTH_TOKEN")
	_ = twilioAccountSid
	_ = twilioAuthToken

	client := twilio.NewRestClient()
	return client
}

func SendMessage(client *twilio.RestClient, message string, receiver string) {
	params := &api.CreateMessageParams{}
	params.SetFrom("whatsapp:" + SENDER)
	params.SetTo("whatsapp:" + receiver)
	params.SetBody(message)

	// TODO: uncomment after debugging
	//resp, err := client.Api.CreateMessage(params)
	//if err != nil {
	//	fmt.Println(err.Error())
	//	os.Exit(1)
	//} else {
	//	if resp.Body != nil {
	//		fmt.Println(*resp.Body)
	//	} else {
	//		fmt.Println(resp.Body)
	//	}
	//}
}

func CreateDailyTaskMessage[T Assignable](tasks []T, member *Member) string {
	message := fmt.Sprintf("%s! Deine heutigen Aufgaben sind:\n", member.Name)

	dailyTasks := "\n"
	for _, task := range tasks {
		dailyTasks += fmt.Sprintf("- %s\n", task.GetName())
	}

	return message + dailyTasks
}
