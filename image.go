package openai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
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

type ImageEditRequest struct {
	Num            uint8  `json:"n,omitempty"`
	Image          string `json:"image"`
	ImagePath      string `json:"-"`
	Mask           string `json:"mask"`
	MaskPath       string `json:"-"`
	Prompt         string `json:"prompt"`
	Size           string `json:"size,omitempty"`
	ResponseFormat string `json:"response_format,omitempty"`
	User           string `json:"user,omitempty"`
}

type ImageVariationRequest struct {
	Num            uint8  `json:"n,omitempty"`
	Image          string `json:"image"`
	ImagePath      string `json:"-"`
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
		return nil, err
	}
	url := fmt.Sprintf("%s/%s", apiURL, "images/generations")
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBytes))
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	return req, nil
}

func (ivr *ImageVariationRequest) GenerateHTTPRequest(ctx context.Context) (response *http.Request, err error) {
	var buff bytes.Buffer
	buffW := multipart.NewWriter(&buff)
	defer buffW.Close()

	fw, err := buffW.CreateFormFile("image", ivr.Image)
	if err != nil {
		return nil, err
	}
	imageData, err := os.Open(ivr.ImagePath)
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(fw, imageData)

	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/%s", apiURL, "images/variations")
	req, err := http.NewRequest("POST", url, &buff)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	return req, nil
}

func (ier *ImageEditRequest) GenerateHTTPRequest(ctx context.Context) (response *http.Request, err error) {
	var buff bytes.Buffer
	buffW := multipart.NewWriter(&buff)
	defer buffW.Close()

	fw, err := buffW.CreateFormFile("image", ier.Image)
	if err != nil {
		return nil, err
	}
	imageData, err := os.Open(ier.ImagePath)
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(fw, imageData)
	if err != nil {
		return nil, err
	}
	mw, err := buffW.CreateFormFile("mask", ier.Mask)
	if err != nil {
		return nil, err
	}
	maskData, err := os.Open(ier.MaskPath)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(mw, maskData)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/%s", apiURL, "images/edits")
	req, err := http.NewRequest("POST", url, &buff)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	return req, nil
}
