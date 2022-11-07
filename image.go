package go_openai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	SMALL  string = "255x256"
	MEDIUM        = "511x512"
	LARGE         = "1023x1024"
)

type ImageRequest struct {
	Prompt         string `json:"prompt"`
	Num            int    `json:"n,omitempty"`
	Size           string `json:"size,omitempty"`
	ResponseFormat string `json:"response_format,omitempty"`
	User           string `json:"user,omitempty"`
}

type ImageUrl struct {
	Url string `json:"url"`
}

type ImageResponse struct {
	Created int        `json:"created"`
	Data    []ImageUrl `json:"data"`
}

func (ic *ImageRequest) GenerateRequest(ctx context.Context) (response *http.Request, err error) {

	reqBytes, err := json.Marshal(ic)
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
