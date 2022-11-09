# go-openai
[OpenAI-GPT3](https://beta.openai.com/) OpenAI API wrapper for Go

Checkout [go-gpt3](https://github.com/sashabaranov/go-gpt3)  API wrapper, If this one does not fit your needs


Currently this wrapper supports the following API's:
- Images
- Models


## Installation:
```
go get github.com/EthanCampana/go-openai 
```

## Example Usage
```go
package main

import (
	"context"
	"fmt"
	openai "github.com/EthanCampana/go-openai"
)

func main() {
	c := openai.GetClient("your token")
	ctx := context.Background()
    rb := openai.GetRequestBuilder("image").(*openai.ImageRequestBuilder)
    res := c.CreateImage(
        ctx ,
        rb.SetPrompt("A Chicken With Glasses, Digtal Art").
            SetNumberOfPictures(3).
            SetSize(oa.LARGE).
            ReturnRequest()
    )
    fmt.Println(res.data[0].url)
}
```
or
```go
package main

import (
	"context"
	"fmt"
	openai "github.com/EthanCampana/go-openai"
)

func main() {
	c := openai.GetClient("your token")
	ctx := context.Background()
    req := &openai.ImageRequest{
        Num:            3,
        Prompt:         "A Chicken With Glasses, Digtal Art",
        Size:           oa.LARGE,
        ResponseFormat: "url",
        User:           "",
        
    }
    res := c.CreateImage(ctx,req)
    fmt.Println(res.data[0].url)
}
```

Both would get the same result. The RequestBuilder struct builds out the requests for you with safeguards in place so that you don't send out a bad requests!

### Request Builders
- image
- image-variation
- image-edit

example usage
```go
    openai.GetRequestBuilder("image").(*openai.ImageRequestBuilder)
    openai.GetRequestBuilder("image-variation").(*openai.ImageVariationRequestBuilder)
    openai.GetRequestBuilder("image-edit").(*openai.ImageEditRequestBuilder)
```
