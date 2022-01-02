package appnubankroutinestartwatchemail

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	".app-nubank-routine-start-watch-email/gateways"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

const (
	_appNubankTopic     = "APP_NUBANK_TOPIC"
	_appNubankMailLabel = "APP_NUBANK_MAIL_LABEL"
)

func StartWatchEmailHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	gmailService, err := gmail.NewService(ctx, option.WithTokenSource(gateways.NewGoogleToken().GetTokenSource()))

	if gateways.HandleError("Create Context", err, w) {
		return
	}

	watchResponse, err := gmailService.Users.Watch("me", createWatchRequest()).Do()

	if gateways.HandleError("Perform API", err, w) {
		return
	}

	json, err := watchResponse.MarshalJSON()

	if gateways.HandleError("Bind", err, w) {
		return
	}

	log.Println(">>", string(json))
	fmt.Fprint(w, string(json))
}

func createWatchRequest() *gmail.WatchRequest {
	return &gmail.WatchRequest{
		LabelFilterAction: "include",
		LabelIds:          []string{os.Getenv(_appNubankMailLabel)},
		TopicName:         os.Getenv(_appNubankTopic),
	}
}
