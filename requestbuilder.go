package openai

import (
	"log"
)

type RequestBuilder interface {
	ReturnRequest() Request
}

func (c *Client) GetRequestBuilder() RequestBuilder {
	imgreq := &ImageRequest{
		Num:            1,
		Prompt:         "",
		Size:           SMALL,
		ResponseFormat: "url",
		User:           "",
	}
	return &ImageRequestBuilder{req: imgreq}
}

const (
	MaxImageRequest = 10
)

type ImageRequestBuilder struct{ req *ImageRequest }

func (irb *ImageRequestBuilder) ReturnRequest() Request {
	return irb.req
}

func (irb *ImageRequestBuilder) SetPrompt(prompt string) *ImageRequestBuilder {
	irb.req.Prompt = prompt
	return irb
}

func (irb *ImageRequestBuilder) SetResponseFormat(rf string) *ImageRequestBuilder {
	switch {
	case rf == "url" || rf == "b64_json":
		irb.req.ResponseFormat = rf
	default:
		log.Println("[WARN] response format you provided is invalid. Setting response format to url")
		irb.req.ResponseFormat = "url"
	}
	return irb
}

func (irb *ImageRequestBuilder) SetUser(user string) *ImageRequestBuilder {
	irb.req.User = user
	return irb
}

func (irb *ImageRequestBuilder) SetNumberOfPictures(num uint8) *ImageRequestBuilder {
	if num > MaxImageRequest {
		log.Println("[WARN] Num you provided is not accepted. Setting Num to 1")
		num = 1
	}
	irb.req.Num = num
	return irb
}

func (irb *ImageRequestBuilder) SetSize(size string) *ImageRequestBuilder {
	switch {
	case size == SMALL:
		irb.req.Size = SMALL
	case size == MEDIUM:
		irb.req.Size = MEDIUM
	case size == LARGE:
		irb.req.Size = LARGE
	default:
		log.Println("[WARN] Size you provided is not accepted. Setting Size to 256x256")
		irb.req.Size = SMALL
	}
	return irb
}
