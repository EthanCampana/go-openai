package openai

type Model struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	OwnedBy string `json:"owned_by"`
}
type Models struct {
	Data []Model `json:"data"`
}
