package openai

import (
	"context"
	"net/http"
)

type Request interface {
	GenerateRequest(context.Context) (*http.Request, error)
}

const apiURL = "https://api.openai.com/v1"

type Client struct {
	authToken  string
	orgID      string
	httpClient *http.Client
}
