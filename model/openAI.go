package model

type OpenAIRequest struct {
	Model     string          `json:"model"`
	Messages  []OpenAIMessage `json:"messages"`
	MaxTokens int             `json:"max_tokens"`
	Stream    bool            `json:"stream"`
}

type OpenAIMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type OpenAIResponse struct {
	ID      string                 `json:"id"`
	Choices []OpenAIResponseChoice `json:"choices"`
}

type OpenAIResponseChoice struct {
	Message OpenAIMessage `json:"message"`
}
