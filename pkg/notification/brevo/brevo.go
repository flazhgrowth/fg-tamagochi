package brevo

import (
	brevolib "github.com/getbrevo/brevo-go/lib"
)

type BrevoImpl struct {
	client *brevolib.APIClient
}

func New(brevoCfg BrevoCfg) *BrevoImpl {
	cfg := brevolib.NewConfiguration()
	cfg.AddDefaultHeader("api-key", brevoCfg.APIKey)
	client := brevolib.NewAPIClient(cfg)

	return &BrevoImpl{
		client: client,
	}
}
