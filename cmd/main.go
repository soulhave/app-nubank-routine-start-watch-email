package main

import (
	"context"
	"log"
	"os"

	app ".app-nubank-routine-start-watch-email"
	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
)

func main() {
	ctx := context.Background()
	if err := funcframework.RegisterHTTPFunctionContext(ctx, "/", app.StartWatchEmailHTTP); err != nil {
		log.Fatalf("app_nubank_routine_start_watch_email.StartWatchEmailHTTP: %v\n", err)
	}
	// Use PORT environment variable, or default to 8080.
	port := "8080"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}
	if err := funcframework.Start(port); err != nil {
		log.Fatalf("funcframework.Start: %v\n", err)
	}
}
