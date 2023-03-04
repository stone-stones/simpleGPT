package chatGPT

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

//defaultConfig is used as a cache to avoid loading the configuration from the file every time.
var defaultConfig ChatConfig

//GetConfig returns a plain configuration with base parameters.
func GetConfig() ChatConfig {
	// If the default configuration has not been initialized yet, initialize it with default values.
	if !defaultConfig.isInit {
		defaultConfig = ChatConfig{
			PostParams: PostParams{
				Model:       DefaultModelName,
				MaxTokens:   2048,
				Temperature: 0.2,
			},
			isInit: true,
		}
	}
	// Make a copy of the default configuration and return it.
	conf := defaultConfig
	return conf
}

//PostParams is a struct that contains the parameters to be sent to the OpenAI API.
type PostParams struct {
	Model            string         `json:"model"`
	Message          []ChatMsg      `json:"messages"`
	MaxTokens        int            `json:"max_tokens,omitempty"`
	Temperature      float32        `json:"temperature,omitempty"`
	TopP             float32        `json:"top_p,omitempty"`
	N                int            `json:"n,omitempty"`
	Stream           bool           `json:"stream,omitempty"`
	Stop             []string       `json:"stop,omitempty"`
	PresencePenalty  float32        `json:"presence_penalty,omitempty"`
	FrequencyPenalty float32        `json:"frequency_penalty,omitempty"`
	LogitBias        map[string]int `json:"logit_bias,omitempty"`
	User             string         `json:"user,omitempty"`
}

//ChatConfig is a struct that contains the configuration for the Chat client.
type ChatConfig struct {
	PostParams PostParams `json:"postParams"`
	Token      string     `json:"token"`
	isInit     bool
}

//LoadConfigFromFile reads the configuration from a file.
func LoadConfigFromFile(path string, log Logger) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Error("ReadFile error:%s", err.Error())
		return err
	}

	if err = json.Unmarshal(file, &defaultConfig); err != nil {
		log.Error("Unmarshal error:%s", err.Error())
		return err
	}

	if defaultConfig.Token == "" ||
		(defaultConfig.PostParams.Model != DefaultModelName &&
			defaultConfig.PostParams.Model != validModelName) {
		err = errors.New("model name or token not valid")
		log.Error(err.Error())
		return err
	}
	defaultConfig.isInit = true
	return nil
}
