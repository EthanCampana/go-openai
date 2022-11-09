package openai

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
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
func (c *Client) setHeaders(r *http.Request) *http.Request {
	r.Header.Set("Accept", "application/json; charset=utf-8")
	r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.authToken))
	if len(c.orgID) > 0 {
		r.Header.Set("OpenAI-Organization", c.orgID)
	}
	return r
}

func checkResponse(resp *http.Response) error {
	defer resp.Body.Close()
	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusBadRequest {
		var errResp APIErrorResponse
		err := json.NewDecoder(resp.Body).Decode(&errResp)
		if err != nil || errResp.Error == nil {
			return fmt.Errorf("error, status code: %d", resp.StatusCode)
		}
		return fmt.Errorf("error, status code: %d, message: %s", resp.StatusCode, errResp.Error.Message)
	}
	return nil
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
	resp, err := c.httpClient.Do(c.setHeaders(req))
	if err != nil {
		return res, err
	}
	if err = checkResponse(resp); err != nil {
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
	err = c.SendRequest(req, res)
	if err != nil {
		return res, err
	}
	return res, nil
}

// Sends an HttpRequest to the OpenAI API and Loads information into the buffer that is passed.
func (c *Client) SendRequest(req *http.Request, a interface{}) error {
	res, err := c.httpClient.Do(c.setHeaders(req))
	if err != nil {
		return err
	}
	if err = checkResponse(res); err != nil {
		return err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(body, &a); err != nil {
		return err
	}
	return nil
}

// Utilizes the CreateImage OpenAI API  to generate Art based on the Request parameters.
//
// @Returns openai.ImageResponse.
func (c *Client) CreateImage(ctx context.Context, imgReq Request) (ImageResponse, error) {
	var imgRes ImageResponse
	var req *http.Request
	var err error
	switch i := imgReq.(type) {
	case *ImageRequest:
		req, err = i.GenerateHTTPRequest(ctx)
	case *ImageVariationRequest:
		req, err = i.GenerateHTTPRequest(ctx)
	case *ImageEditRequest:
		req, err = i.GenerateHTTPRequest(ctx)
	default:
		return imgRes, fmt.Errorf("got unsupported request type %T", imgReq)
	}
	if err != nil {
		return imgRes, err
	}
	err = c.SendRequest(req, imgRes)
	if err != nil {
		return imgRes, err
	}
	return imgRes, err
}

// Utilizes the CreateImageVariation OpenAI API to generate Art based on the Request parameters.
//
// @Returns openai.ImageResponse.
func (c *Client) CreateImageVariation(ctx context.Context, imgVarReq *ImageVariationRequest) (ImageResponse, error) {
	return c.CreateImage(ctx, imgVarReq)
}

// Utilizes the CreateImageEdit OpenAI API to generate Art based on the Request parameters.
//
// @Returns openai.ImageResponse.
func (c *Client) CreateImageEidt(ctx context.Context, imgEditReq *ImageEditRequest) (ImageResponse, error) {
	return c.CreateImage(ctx, imgEditReq)
}
