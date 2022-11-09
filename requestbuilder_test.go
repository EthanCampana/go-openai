package openai_test

import (
	"reflect"
	"testing"

	. "github.com/EthanCampana/go-openai"
)

func TestImageRequestBuilder_ReturnRequest(t *testing.T) {
	type fields struct {
		Req *ImageRequest
	}
	tests := []struct {
		name   string
		fields fields
		want   Request
	}{
		{
			name: "Return Basic Request",
			fields: fields{
				Req: &ImageRequest{
					Num:            3,
					Prompt:         "Hello World",
					Size:           "256x256",
					ResponseFormat: "",
					User:           "",
				},
			},
			want: &ImageRequest{
				Num:            3,
				Prompt:         "Hello World",
				Size:           "256x256",
				ResponseFormat: "",
				User:           "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			irb := ImageRequestBuilder{
				Req: tt.fields.Req,
			}
			if got := irb.ReturnRequest(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ImageRequestBuilder.ReturnRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestImageVariationRequestBuilder_ReturnRequest(t *testing.T) {
	c := GetClient("Some Auth Token")
	irb, _ := c.GetRequestBuilder("image-variation").(ImageVariationRequestBuilder)
	type fields struct {
		Irb       ImageRequestBuilder
		Image     string
		ImagePath string
	}
	tests := []struct {
		name   string
		fields fields
		want   Request
	}{
		{
			name: "Return Basic ImageVariationRequest",
			fields: fields{
				Irb:       irb.Irb,
				Image:     "",
				ImagePath: "",
			},
			want: &ImageVariationRequest{
				Num:            1,
				Prompt:         "",
				Size:           SMALL,
				ResponseFormat: "url",
				User:           "",
				Image:          "test.ty",
				ImagePath:      "home/test.ty",
			},
		},
		{
			name: "Check File Path in Same Directory",
			fields: fields{
				Irb:       irb.Irb,
				Image:     "",
				ImagePath: "",
			},
			want: &ImageVariationRequest{
				Num:            1,
				Prompt:         "",
				Size:           SMALL,
				ResponseFormat: "url",
				User:           "",
				Image:          "test.ty",
				ImagePath:      "./test.ty",
			},
		},
		{
			name: "Check File Path in Same Directory without",
			fields: fields{
				Irb:       irb.Irb,
				Image:     "",
				ImagePath: "",
			},
			want: &ImageVariationRequest{
				Num:            1,
				Prompt:         "",
				Size:           SMALL,
				ResponseFormat: "url",
				User:           "",
				Image:          "test.ty",
				ImagePath:      "test.ty",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ivrb := ImageVariationRequestBuilder{
				Irb:       tt.fields.Irb,
				Image:     tt.fields.Image,
				ImagePath: tt.fields.ImagePath,
			}
			ivrb.SetImage(tt.want.(*ImageVariationRequest).ImagePath)
			if got := ivrb.ReturnRequest(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ImageVariationRequestBuilder.ReturnRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuildingEditRequest(t *testing.T) {
	c := GetClient("Some Auth Token")
	testIerb, _ := c.GetRequestBuilder("image-edit").(ImageEditRequestBuilder)
	type args struct {
		Num            uint8
		Prompt         string
		Size           string
		ResponseFormat string
		User           string
		ImagePath      string
		MaskPath       string
	}
	tests := []struct {
		name string
		ierb ImageEditRequestBuilder
		args args
		want Request
	}{
		{
			name: "Building Edit Request from scratch",
			ierb: testIerb,
			args: args{
				Num:            10,
				Prompt:         "Testing",
				Size:           "256x256",
				ResponseFormat: "url",
				User:           "",
				ImagePath:      "testing/test",
				MaskPath:       "testing/mask",
			},
			want: &ImageEditRequest{
				Num:            10,
				Prompt:         "Testing",
				Size:           "256x256",
				ResponseFormat: "url",
				User:           "",
				ImagePath:      "testing/test",
				Image:          "test",
				Mask:           "mask",
				MaskPath:       "testing/mask",
			},
		},
		{
			name: "Building Bad Edit Request",
			ierb: testIerb,
			args: args{
				Num:            112,
				Prompt:         "Testing",
				Size:           "1024x256",
				ResponseFormat: "json",
				User:           "6657",
				ImagePath:      "testing/test",
				MaskPath:       "testing/mask",
			},
			want: &ImageEditRequest{
				Num:            1,
				Prompt:         "Testing",
				Size:           "256x256",
				ResponseFormat: "url",
				User:           "6657",
				ImagePath:      "testing/test",
				Image:          "test",
				Mask:           "mask",
				MaskPath:       "testing/mask",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ierb.SetNumberOfPictures(tt.args.Num).
				SetSize(tt.args.Size).
				SetPrompt(tt.args.Prompt).
				SetResponseFormat(tt.args.ResponseFormat).
				SetUser(tt.args.User).
				SetImage(tt.args.ImagePath).
				SetMask(tt.args.MaskPath)

			if got := tt.ierb.ReturnRequest(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ImageEditRequestBuilder.ReturnRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestImageEditRequestBuilder_ReturnRequest(t *testing.T) {
	c := GetClient("Some Auth Token")
	testIerb, _ := c.GetRequestBuilder("image-edit").(ImageEditRequestBuilder)
	type fields struct {
		Irb       ImageRequestBuilder
		Image     string
		ImagePath string
		Mask      string
		MaskPath  string
	}
	tests := []struct {
		name   string
		fields fields
		want   Request
	}{
		{
			name: "Generate Basic Image Edit Request",
			fields: fields{
				Irb:       testIerb.Irb,
				Image:     testIerb.Image,
				ImagePath: testIerb.ImagePath,
				Mask:      testIerb.Mask,
				MaskPath:  testIerb.MaskPath,
			},
			want: &ImageEditRequest{
				Num:            1,
				Prompt:         "",
				Size:           SMALL,
				ResponseFormat: "url",
				User:           "",
				Image:          "test.ty",
				ImagePath:      "home/test.ty",
				Mask:           "mask.ty",
				MaskPath:       "Masks/mask.ty",
			},
		},
		{
			name: "Check Mask Path in Same Directory",
			fields: fields{
				Irb:       testIerb.Irb,
				Image:     testIerb.Image,
				ImagePath: testIerb.ImagePath,
				Mask:      testIerb.Mask,
				MaskPath:  testIerb.MaskPath,
			},
			want: &ImageEditRequest{
				Num:            1,
				Prompt:         "",
				Size:           SMALL,
				ResponseFormat: "url",
				User:           "",
				Image:          "test.ty",
				ImagePath:      "home/test.ty",
				Mask:           "mask.ty",
				MaskPath:       "/mask.ty",
			},
		},
		{
			name: "Ensure Paths Are Set",
			fields: fields{
				Irb:       testIerb.Irb,
				Image:     testIerb.Image,
				ImagePath: testIerb.ImagePath,
				Mask:      testIerb.Mask,
				MaskPath:  testIerb.MaskPath,
			},
			want: &ImageEditRequest{
				Num:            1,
				Prompt:         "",
				Size:           SMALL,
				ResponseFormat: "url",
				User:           "",
				Image:          "test.ty",
				ImagePath:      "test.ty",
				Mask:           "mask.ty",
				MaskPath:       "./mask.ty",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ierb := ImageEditRequestBuilder{
				Irb:       tt.fields.Irb,
				Image:     tt.fields.Image,
				ImagePath: tt.fields.ImagePath,
				Mask:      tt.fields.Mask,
				MaskPath:  tt.fields.MaskPath,
			}
			ierb.SetImage(tt.want.(*ImageEditRequest).ImagePath)
			ierb.SetMask(tt.want.(*ImageEditRequest).MaskPath)
			if got := ierb.ReturnRequest(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ImageEditRequestBuilder.ReturnRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetRequestBuilder(t *testing.T) {
	type args struct {
		builder string
	}
	tests := []struct {
		name string
		args args
		want RequestBuilder
	}{
		{
			name: "Get ImageRequestBuilder Struct",
			args: args{builder: "image"},
			want: ImageRequestBuilder{Req: &ImageRequest{
				Num:            1,
				Prompt:         "",
				Size:           SMALL,
				ResponseFormat: "url",
				User:           "",
			}},
		},
		{
			name: "Get ImageVariantionRequestBuilder Struct",
			args: args{builder: "image-variation"},
			want: ImageVariationRequestBuilder{
				Irb: ImageRequestBuilder{
					Req: &ImageRequest{
						Num:            1,
						Prompt:         "",
						Size:           SMALL,
						ResponseFormat: "url",
						User:           "",
					},
				},
				Image:     "",
				ImagePath: "",
			},
		},
		{
			name: "Get ImageEditRequestBuilder Struct",
			args: args{builder: "image-edit"},
			want: ImageEditRequestBuilder{
				Irb: ImageRequestBuilder{
					Req: &ImageRequest{
						Num:            1,
						Prompt:         "",
						Size:           SMALL,
						ResponseFormat: "url",
						User:           "",
					},
				},
				Image:     "",
				ImagePath: "",
				Mask:      "",
				MaskPath:  "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := GetClient("SOME TOKEN")
			if got := c.GetRequestBuilder(tt.args.builder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.GetRequestBuilder() = %v, want %v", got, tt.want)
			}
		})
	}
}
