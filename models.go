package openai

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Model struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	OwnedBy string `json:"owned_by"`
}
type Models struct {
	Data []Model `json:"data"`
}

// Calls OpenAI GetModel API to gather Model information of the Model Id provided
//
// @Returns openai.Model Struct.
func (c *Client) GetModel(ctx context.Context, model string) (Model, error) {
	var res Model
	url := fmt.Sprintf("https://api.openai.com/v1/models/%s", model)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return res, err
	}
	req = req.WithContext(ctx)
	req.Header.Set("Accept", "application/json; charset=utf-8")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.authToken))
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	if err = json.Unmarshal(body, &res); err != nil {
		return res, err
	}
	return res, nil
}

// Calls OpenAI ListModel API. Provides A list of Models currently supported by OpenAI
//
// @Returns openai.Models Struct.
func (c *Client) ListModels(ctx context.Context) (Models, error) {
	var res Models
	req, err := http.NewRequest("GET", "https://api.openai.com/v1/models", nil)
	if err != nil {
		return res, err
	}
	req = req.WithContext(ctx)
	req.Header.Set("Accept", "application/json; charset=utf-8")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.authToken))
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	if err = json.Unmarshal(body, &res); err != nil {
		return res, err
	}
	return res, nil
}
