package api

type OpenAIRequest struct {
	Model       string       `json:"model"`
	Messages    []OAIMessage `json:"messages"`
	MaxToken    int          `json:"max_tokens,omitempty"`
	Temperature *float64     `json:"temperature,omitempty"`
	Stream      bool         `json:"stream,omitempty"`
}

type OAIMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type OpenAIResponse struct {
	ID      string      `json:"id"`
	Model   string      `json:"model"`
	Choices []OAIChoice `json:"choices"`
	Usage   *OAIUsage   `json:"usage"`
}

type OAIChoice struct {
	Index        string     `json:"index"`
	Message      OAIMessage `json:"message"`
	FinishReason string     `json:"finish_reason"`
}

type OAIUsage struct {
	PromptsToken     int `json:"prompts_token"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}
