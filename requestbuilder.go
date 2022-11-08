package openai

import (
	"log"
	"strings"
)

type RequestBuilder interface {
	ReturnRequest() Request
}

func (c *Client) GetRequestBuilder(builder string) RequestBuilder {
	var b RequestBuilder
	switch {
	case builder == "image":
		imgreq := &ImageRequest{
			Num:            1,
			Prompt:         "",
			Size:           SMALL,
			ResponseFormat: "url",
			User:           "",
		}
		b = &ImageRequestBuilder{req: imgreq}
	case builder == "image-variation":
		irb, _ := c.GetRequestBuilder("image").(ImageRequestBuilder)
		req := imageRequestToImageVariationRequest(irb.req)
		b = &ImageVariationRequestBuilder{
			irb: irb,
			req: req,
		}
	}
	return b
}

func imageRequestToImageVariationRequest(ir *ImageRequest) *ImageVariationRequest {
	return &ImageVariationRequest{
		Num:       ir.Num,
		Prompt:    ir.Prompt,
		Size:      ir.Size,
		User:      ir.User,
		Image:     "",
		ImagePath: "",
	}
}

const (
	MaxImageRequest = 10
)

type ImageRequestBuilder struct{ req *ImageRequest }

type ImageVariationRequestBuilder struct {
	irb ImageRequestBuilder
	req *ImageVariationRequest
}

// type ImageEditRequestBuilder struct {
// 	irb ImageVariationRequestBuilder
// 	req *ImageEditRequest
// }

func (ivrb ImageVariationRequestBuilder) ReturnRequest() Request {
	ivrb.req.Num = ivrb.irb.req.Num
	ivrb.req.Prompt = ivrb.irb.req.Prompt
	ivrb.req.Size = ivrb.irb.req.Size
	ivrb.req.ResponseFormat = ivrb.irb.req.ResponseFormat
	ivrb.req.User = ivrb.irb.req.User
	return ivrb.req
}

func (ivrb *ImageVariationRequestBuilder) SetImage(filepath string) *ImageVariationRequestBuilder {
	ivrb.req.ImagePath = filepath
	ivrb.req.Image = strings.SplitAfter(filepath, "/")[len(strings.SplitAfter(filepath, "/"))-1]
	return ivrb
}

func (irb ImageRequestBuilder) ReturnRequest() Request {
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

func (ivrb *ImageVariationRequestBuilder) SetSize(size string) *ImageVariationRequestBuilder {
	ivrb.irb.SetSize(size)
	return ivrb
}
func (ivrb *ImageVariationRequestBuilder) SetNumberOfPictures(num uint8) *ImageVariationRequestBuilder {
	ivrb.irb.SetNumberOfPictures(num)
	return ivrb
}
func (ivrb *ImageVariationRequestBuilder) SetUser(user string) *ImageVariationRequestBuilder {
	ivrb.irb.SetUser(user)
	return ivrb
}
func (ivrb *ImageVariationRequestBuilder) SetResponseFormat(rf string) *ImageVariationRequestBuilder {
	ivrb.irb.SetResponseFormat(rf)
	return ivrb
}
func (ivrb *ImageVariationRequestBuilder) SetPrompt(prompt string) *ImageVariationRequestBuilder {
	ivrb.irb.SetPrompt(prompt)
	return ivrb
}
