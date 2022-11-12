package openai

import (
	"fmt"
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
		b = ImageRequestBuilder{Req: imgreq}
	case builder == "image-variation":
		irb, _ := c.GetRequestBuilder("image").(ImageRequestBuilder)
		b = ImageVariationRequestBuilder{
			Irb:       irb,
			Image:     "",
			ImagePath: "",
		}
	case builder == "image-edit":
		irb, _ := c.GetRequestBuilder("image").(ImageRequestBuilder)
		b = ImageEditRequestBuilder{
			Irb:       irb,
			Image:     "",
			ImagePath: "",
			Mask:      "",
			MaskPath:  "",
		}
	default:
		err := fmt.Errorf("%s is not an accepted builder option", builder)
		log.Printf("Function Call GetRequestBuilder failed with builder: %s %v", builder, err)
	}
	return b
}

func imageRequestToImageVariationRequest(ir *ImageRequest) *ImageVariationRequest {
	return &ImageVariationRequest{
		Num:            ir.Num,
		Prompt:         ir.Prompt,
		Size:           ir.Size,
		User:           ir.User,
		ResponseFormat: ir.ResponseFormat,
		Image:          "",
		ImagePath:      "",
	}
}

func imageRequestToImageEditRequest(ivr *ImageRequest) *ImageEditRequest {
	return &ImageEditRequest{
		Num:            ivr.Num,
		Prompt:         ivr.Prompt,
		Size:           ivr.Size,
		User:           ivr.User,
		ResponseFormat: ivr.ResponseFormat,
		Image:          "",
		ImagePath:      "",
		Mask:           "",
		MaskPath:       "",
	}
}

const (
	MaxImageRequest = 10
)

type ImageRequestBuilder struct{ Req *ImageRequest }

type ImageVariationRequestBuilder struct {
	Irb       ImageRequestBuilder
	Image     string
	ImagePath string
}

type ImageEditRequestBuilder struct {
	Irb       ImageRequestBuilder
	Image     string
	ImagePath string
	Mask      string
	MaskPath  string
}

func (ierb ImageEditRequestBuilder) ReturnRequest() Request {
	ier := imageRequestToImageEditRequest(ierb.Irb.Req)
	ier.Mask = ierb.Mask
	ier.MaskPath = ierb.MaskPath
	ier.Image = ierb.Image
	ier.ImagePath = ierb.ImagePath
	if ier.Image == "" || ier.Mask == "" {
		log.Fatal("Cannot Generate Request if ImagePath or MaskPath is empty!")
	}
	return ier
}

// Returns the Underlying Request of the Given RequestBuilder.
func (ivrb ImageVariationRequestBuilder) ReturnRequest() Request {
	ivr := imageRequestToImageVariationRequest(ivrb.Irb.Req)
	ivr.Image = ivrb.Image
	ivr.ImagePath = ivrb.ImagePath
	if ivr.Image == "" {
		log.Fatal("Cannot Generate Request if ImagePath is empty!")
	}
	return ivr
}

// Sets the image to upload to upload of the underlying Request.
func (ivrb *ImageVariationRequestBuilder) SetImage(filepath string) *ImageVariationRequestBuilder {
	ivrb.ImagePath = filepath
	ivrb.Image = strings.SplitAfter(filepath, "/")[len(strings.SplitAfter(filepath, "/"))-1]
	return ivrb
}

// Returns the Underlying Request of the Given RequestBuilder.
func (irb ImageRequestBuilder) ReturnRequest() Request {
	return irb.Req
}

// Sets the prompt of the underlying Request.
func (irb *ImageRequestBuilder) SetPrompt(prompt string) *ImageRequestBuilder {
	irb.Req.Prompt = prompt
	return irb
}

// Sets the ResponseFormat of the underlying Request.
func (irb *ImageRequestBuilder) SetResponseFormat(rf string) *ImageRequestBuilder {
	switch {
	case rf == "url" || rf == "b64_json":
		irb.Req.ResponseFormat = rf
	default:
		log.Println("[WARN] response format you provided is invalid. Setting response format to url")
		irb.Req.ResponseFormat = "url"
	}
	return irb
}

// Sets the user of the underlying Request.
func (irb *ImageRequestBuilder) SetUser(user string) *ImageRequestBuilder {
	irb.Req.User = user
	return irb
}

// Sets the number of images to generate of the underlying Request.
//
// min=1, max=10, default=1.
func (irb *ImageRequestBuilder) SetNumberOfPictures(num uint8) *ImageRequestBuilder {
	if num > MaxImageRequest {
		log.Println("[WARN] Num you provided is not accepted. Setting Num to 1")
		num = uint8(1)
	}
	irb.Req.Num = num
	return irb
}

// Sets the size of the underlying Request
//
// SMALL = 256x256  MEDIUM = 512x512  LARGE = 1024x1024  Default = 256x265.
func (irb *ImageRequestBuilder) SetSize(size string) *ImageRequestBuilder {
	switch {
	case size == SMALL:
		irb.Req.Size = SMALL
	case size == MEDIUM:
		irb.Req.Size = MEDIUM
	case size == LARGE:
		irb.Req.Size = LARGE
	default:
		log.Println("[WARN] Size you provided is not accepted. Setting Size to 256x256")
		irb.Req.Size = SMALL
	}
	return irb
}

// Sets the size of the underlying Request
//
// SMALL = 256x256  MEDIUM = 512x512  LARGE = 1024x1024  Default = 256x265.
func (ivrb *ImageVariationRequestBuilder) SetSize(size string) *ImageVariationRequestBuilder {
	ivrb.Irb.SetSize(size)
	return ivrb
}

// Sets the number of images to generate of the underlying Request.
//
// min=1, max=10, default=1.
func (ivrb *ImageVariationRequestBuilder) SetNumberOfPictures(num uint8) *ImageVariationRequestBuilder {
	ivrb.Irb.SetNumberOfPictures(num)
	return ivrb
}

// Sets the user of the underlying Request.
func (ivrb *ImageVariationRequestBuilder) SetUser(user string) *ImageVariationRequestBuilder {
	ivrb.Irb.SetUser(user)
	return ivrb
}

// Sets the ResponseFormat of the underlying Request.
func (ivrb *ImageVariationRequestBuilder) SetResponseFormat(rf string) *ImageVariationRequestBuilder {
	ivrb.Irb.SetResponseFormat(rf)
	return ivrb
}

// Sets the prompt of the underlying Request.
func (ivrb *ImageVariationRequestBuilder) SetPrompt(prompt string) *ImageVariationRequestBuilder {
	ivrb.Irb.SetPrompt(prompt)
	return ivrb
}

// Sets the size of the underlying Request
//
// SMALL = 256x256  MEDIUM = 512x512  LARGE = 1024x1024  Default = 256x265.
func (ierb *ImageEditRequestBuilder) SetSize(size string) *ImageEditRequestBuilder {
	ierb.Irb.SetSize(size)
	return ierb
}

// Sets the number of images to generate of the underlying Request.
//
// min=1, max=10, default=1.
func (ierb *ImageEditRequestBuilder) SetNumberOfPictures(num uint8) *ImageEditRequestBuilder {
	ierb.Irb.SetNumberOfPictures(num)
	return ierb
}

// Sets the user of the underlying Request.
func (ierb *ImageEditRequestBuilder) SetUser(user string) *ImageEditRequestBuilder {
	ierb.Irb.SetUser(user)
	return ierb
}

// Sets the ResponseFormat of the underlying Request.
func (ierb *ImageEditRequestBuilder) SetResponseFormat(rf string) *ImageEditRequestBuilder {
	ierb.Irb.SetResponseFormat(rf)
	return ierb
}

// Sets the prompt of the underlying Request.
func (ierb *ImageEditRequestBuilder) SetPrompt(prompt string) *ImageEditRequestBuilder {
	ierb.Irb.SetPrompt(prompt)
	return ierb
}

// Sets the image to upload to upload of the underlying Request.
func (ierb *ImageEditRequestBuilder) SetImage(filepath string) *ImageEditRequestBuilder {
	ierb.ImagePath = filepath
	ierb.Image = strings.SplitAfter(filepath, "/")[len(strings.SplitAfter(filepath, "/"))-1]
	return ierb
}

// Sets the mask image to upload of the underlying Request.
func (ierb *ImageEditRequestBuilder) SetMask(filepath string) *ImageEditRequestBuilder {
	ierb.MaskPath = filepath
	ierb.Mask = strings.SplitAfter(filepath, "/")[len(strings.SplitAfter(filepath, "/"))-1]
	return ierb
}
