package chatGPT

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

//defaultConfig for cache base config ,avoid load from file every time
var defaultConfig ChatConfig

//GetConfig get a plain config with base params
func GetConfig() ChatConfig {
	if !defaultConfig.isInit {
		defaultConfig = ChatConfig{
			PostParams: PostParams {
				Model: DefaultModelName,
				MaxTokens:2048,
				Temperature: 0.2,
			},
			isInit: true,
		}
	}
	conf := defaultConfig
	return conf

}

type PostParams struct {
	Model            string         `json:"model"`
	Message 		  []ChatMsg     `json:"messages"`
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


type ChatConfig struct {
	PostParams PostParams `json:"PostParams"`
	Token      string `json:"token"`
	isInit     bool
}

//LoadConfigFromFile read config from file,
func LoadConfigFromFile(path string,log Logger) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Error("ReadFile error:%s",err.Error())
		return err
	}

	if err = json.Unmarshal(file, &defaultConfig); err != nil {
		log.Error("Unmarshal error:%s",err.Error())
		return err
	}

	if defaultConfig.Token == "" ||
		(defaultConfig.PostParams.Model != DefaultModelName &&
		defaultConfig.PostParams.Model != validModelName)  {
		err = errors.New("model name or token  not valid")
		log.Error(err.Error())
		return err
	}
	defaultConfig.isInit =true
	return nil
}