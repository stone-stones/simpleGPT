package chatGPT

const (
	//StreamEndSign stream response end sign
	StreamEndSign = "[DONE]"
	//GPTURL chatGPT api path
	GPTURL = "https://api.openai.com/v1/chat/completions"
	//DefaultModelName model name
	DefaultModelName = "gpt-3.5-turbo"
	//validModelName another valid model name
	validModelName = "gpt-3.5-turbo-0301"
	//RecvChanSize chan to hand stream response messages
	RecvChanSize = 100
)

const (
	//RoleUser user
	RoleUser      = "user"
	RoleSystem    = "system"
	RoleAssistant = "assistant"
)
