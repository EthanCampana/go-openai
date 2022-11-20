package openai

type CompletionRequest[T string | []string] struct {
	LogitBias        map[string]int16 `json:"logit_bias,omitempty"`
	Model            string           `json:"model"`
	Prompt           T                `json:"prompt"`
	Suffix           string           `json:"suffix,omitempty"`
	Stop             T                `json:"stop,omitempty"`
	User             string           `json:"user,omitempty"`
	Temperature      float32          `json:"temperature"`
	TopP             float32          `json:"top_p"`
	FrequencePenalty float32          `json:"frequency_penalty"`
	PresencePenalty  float32          `json:"presence_penalty"`
	MaxTokens        uint8            `json:"max_tokens"`
	Logprobs         uint8            `json:"logprobs,omitempty"`
	N                uint8            `json:"n"`
	BestOf           uint8            `json:"best_of"`
	Stream           bool             `json:"stream"`
	Echo             bool             `json:"echo"`
}

type ChoiceResponse struct {
	Index        uint16   `json:"index"`
	Text         string   `json:"text"`
	Logprobs     []string `json:"logprobs,omitempty"`
	FinishReason string   `json:"finish_reason"`
}

type CompletionResponse struct {
	ID      string           `json:"id"`
	Object  string           `json:"object"`
	Created int              `json:"created"`
	Model   string           `json:"model"`
	Choices []ChoiceResponse `json:"choices"`
	Usage   Usage            `json:"usage"`
}
