package chatGPT

//BaseResponse common response
type BaseResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int64  `json:"created"`
	Model   string `json:"model"`
}

//StreamResponse stream request response
type StreamResponse struct {
	BaseResponse
	Choices []StreamChoice `json:"choices"`
}

//StreamChoice message in stream response
type StreamChoice struct {
	Index        int    `json:"index"`
	FinishReason string `json:"finish_reason"`
	Delta        Delta  `json:"delta"`
}

//Delta info in stream message
type Delta struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

//ChatMsg messages
type ChatMsg struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

//HttpResponse normal response
type HttpResponse struct {
	BaseResponse
	Choices []HttpChoice `json:"choices"`
	Usage   Usage        `json:"usage"`
}

// Usage Represents the total token usage per request to OpenAI.
type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

//HttpChoice normal response message
type HttpChoice struct {
	Index        int     `json:"index"`
	FinishReason string  `json:"finish_reason"`
	Messages     ChatMsg `json:"message"`
}

// APIError provides error information returned by the OpenAI API.
type APIError struct {
	Code       *string `json:"code,omitempty"`
	Message    string  `json:"message"`
	Param      *string `json:"param,omitempty"`
	Type       string  `json:"type"`
	StatusCode int     `json:"-"`
}

// RequestError provides information about generic request errors.
type RequestError struct {
	StatusCode int
	Err        error
}

type ErrorResponse struct {
	Error *APIError `json:"error,omitempty"`
}

func (e *APIError) Error() string {
	return e.Message
}
