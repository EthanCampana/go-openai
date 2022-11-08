package openai

import (
	"log"
	"strings"
)

type RequestBuilder interface {
	ReturnRequest() Request
}

/*
Creates A RequestBuilder and Returns it.

Params: builder -> Name of the builder you would like to Return

RequestBuilders:

- image

- image-variation

- image-edit
.
*/
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

// Returns the Underlying Request of the Given RequestBuilder.
func (ivrb ImageVariationRequestBuilder) ReturnRequest() Request {
	ivr := imageRequestToImageVariationRequest(ivrb.irb.req)
	ivr.Image = ivrb.req.Image
	ivr.ImagePath = ivrb.req.ImagePath
	ivrb.req = ivr
	return ivrb.req
}

// Sets the image to upload to upload of the underlying Request.
func (ivrb *ImageVariationRequestBuilder) SetImage(filepath string) *ImageVariationRequestBuilder {
	ivrb.req.ImagePath = filepath
	ivrb.req.Image = strings.SplitAfter(filepath, "/")[len(strings.SplitAfter(filepath, "/"))-1]
	return ivrb
}

// Returns the Underlying Request of the Given RequestBuilder.
func (irb ImageRequestBuilder) ReturnRequest() Request {
	return irb.req
}

// Sets the prompt of the underyling Request.
func (irb *ImageRequestBuilder) SetPrompt(prompt string) *ImageRequestBuilder {
	irb.req.Prompt = prompt
	return irb
}

// Sets the ResponseFormat of the underlying Request.
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

// Sets the user of the underlying Request.
func (irb *ImageRequestBuilder) SetUser(user string) *ImageRequestBuilder {
	irb.req.User = user
	return irb
}

// Sets the number of images to generate of the underlying Request.
//
// min=1, max=10, default=1.
func (irb *ImageRequestBuilder) SetNumberOfPictures(num uint8) *ImageRequestBuilder {
	if num > MaxImageRequest {
		log.Println("[WARN] Num you provided is not accepted. Setting Num to 1")
		num = 1
	}
	irb.req.Num = num
	return irb
}

// Sets the size of the underlying Request
//
// SMALL = 256x256  MEDIUM = 512x512  LARGE = 1024x1024  Default = 256x265.
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

// Sets the size of the underlying Request
//
// SMALL = 256x256  MEDIUM = 512x512  LARGE = 1024x1024  Default = 256x265.
func (ivrb *ImageVariationRequestBuilder) SetSize(size string) *ImageVariationRequestBuilder {
	ivrb.irb.SetSize(size)
	return ivrb
}

// Sets the number of images to generate of the underlying Request.
//
// min=1, max=10, default=1.
func (ivrb *ImageVariationRequestBuilder) SetNumberOfPictures(num uint8) *ImageVariationRequestBuilder {
	ivrb.irb.SetNumberOfPictures(num)
	return ivrb
}

// Sets the user of the underlying Request.
func (ivrb *ImageVariationRequestBuilder) SetUser(user string) *ImageVariationRequestBuilder {
	ivrb.irb.SetUser(user)
	return ivrb
}

// Sets the ResponseFormat of the underlying Request.
func (ivrb *ImageVariationRequestBuilder) SetResponseFormat(rf string) *ImageVariationRequestBuilder {
	ivrb.irb.SetResponseFormat(rf)
	return ivrb
}

// Sets the prompt of the underyling Request.
func (ivrb *ImageVariationRequestBuilder) SetPrompt(prompt string) *ImageVariationRequestBuilder {
	ivrb.irb.SetPrompt(prompt)
	return ivrb
}
