package chatGPT

const (
	StreamEndSign ="[DONE]"
	GPTURL = "https://api.openai.com/v1/chat/completions"
	DefaultModelName ="gpt-3.5-turbo"
	validModelName = "gpt-3.5-turbo-0301"
)


const (
	RoleUser = "user"
	RoleSystem= "system"
	RoleAssistant = "assistant"
)