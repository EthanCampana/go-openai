package openai

type Usage struct {
	PromptTokens     int `json:"prompt_token,omitempty"`
	CompletionTokens int `json:"completion_tokens,omitempty"`
	TotalTokens      int `json:"total_tokens"`
}
