package openai

import (
	"context"
	"net/http"
	"time"
)

type Request interface {
	GenerateHTTPRequest(context.Context) (*http.Request, error)
}

const apiURL = "https://api.openai.com/v1"

type Client struct {
	authToken  string
	orgID      string
	httpClient *http.Client
}

func getTransportClient() *http.Client {
	return &http.Client{
		Timeout: time.Minute,
	}
}

func GetClient(authToken string) *Client {
	return &Client{authToken: authToken,
		orgID:      "",
		httpClient: getTransportClient(),
	}
}

func GetOrgClient(authToken string, orgID string) *Client {
	return &Client{authToken: authToken,
		orgID:      orgID,
		httpClient: getTransportClient(),
	}
}
