package openai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	SMALL  string = "256x256"
	MEDIUM string = "512x512"
	LARGE  string = "1024x1024"
)

type ImageRequest struct {
	Num            uint8  `json:"n,omitempty"`
	Prompt         string `json:"prompt"`
	Size           string `json:"size,omitempty"`
	ResponseFormat string `json:"response_format,omitempty"`
	User           string `json:"user,omitempty"`
}

type ImageURL struct {
	URL string `json:"url"`
}

type ImageResponse struct {
	Created int        `json:"created"`
	Data    []ImageURL `json:"data"`
}

func (ir *ImageRequest) GenerateHTTPRequest(ctx context.Context) (response *http.Request, err error) {
	reqBytes, err := json.Marshal(ir)
	if err != nil {
		return
	}
	url := fmt.Sprintf("%s/%s", apiURL, "images/generations")
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBytes))
	if err != nil {
		return
	}
	req = req.WithContext(ctx)
	return req, nil
}
