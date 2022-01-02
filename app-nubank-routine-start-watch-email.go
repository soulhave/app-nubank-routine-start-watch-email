package appnubankroutinestartwatchemail

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"golang.org/x/oauth2"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

func StartWatchEmailHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	token := new(oauth2.Token)
	token.AccessToken = ""
	token.RefreshToken = ""
	token.Expiry = time.Time{}
	token.TokenType = "Bearer"

	config := &oauth2.Config{}

	gmailService, err := gmail.NewService(ctx, option.WithTokenSource(config.TokenSource(ctx, token)))

	if handleError("Create Context", err, w) {
		return
	}

	watchRequest := createWatchRequest()

	watchResponse, err := gmailService.Users.Watch("", &watchRequest).Do()

	if handleError("Perform API", err, w) {
		return
	}

	json, err := watchResponse.MarshalJSON()

	if handleError("Bind", err, w) {
		return
	}

	log.Println(">>", string(json))
	fmt.Fprint(w, string(json))
}

func createWatchRequest() gmail.WatchRequest {
	return gmail.WatchRequest{
		LabelFilterAction: "include",
		LabelIds:          []string{"Label1782046973960748351"},
		TopicName:         "",
	}
}

func handleError(phase string, err error, w http.ResponseWriter) bool {
	if err != nil {
		log.Println("#######", phase, "Error:", err)
		fmt.Fprint(w, "{error:\"", err, "\"}")
		return true
	}
	return false
}
