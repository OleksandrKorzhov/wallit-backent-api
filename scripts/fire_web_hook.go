//go:build ignore

package main

import (
	"context"
	"github.com/plaid/plaid-go/v10/plaid"
	"log"
	"wallit/internal/config"
	plaidService "wallit/internal/plaid"
)

func main() {
	config.Init()
	plaidClient := plaidService.New().GetClient()

	request := plaid.NewSandboxItemFireWebhookRequest(
		"access-sandbox-3e7127f8-edfd-4005-b65a-c6cad3c94a9f",
		"SYNC_UPDATES_AVAILABLE",
	)
	_, _, err := plaidClient.PlaidApi.SandboxItemFireWebhook(context.Background()).
		SandboxItemFireWebhookRequest(*request).
		Execute()

	if err != nil {
		log.Fatalf("triggering plaid webhook: %v", err)
	}
}
