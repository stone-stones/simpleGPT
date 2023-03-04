package chatGPT

import (
	"context"
	"encoding/json"
	"net/http"
	"sync"
	"time"
)

// Client represents a client that can send messages to a chatGPT server
type Client struct {
	User          string          // User represents the user of the client
	token         string          // token represents the token used for authentication with the chatGPT server
	stream        bool            // stream is a boolean flag indicating whether streaming is enabled
	MaxPreMessage int             // MaxPreMessage represents the maximum number of messages that can be stored in preMessage
	CreateTime    int64           // CreateTime represents the time the client was created
	LastConnTime  int64           // LastConnTime represents the time the client last connected to the server
	preMessage    []ChatMsg       // preMessage is a slice containing previously sent messages
	Ctx           context.Context // Ctx represents the context for the client
	mutex         *sync.Mutex     // mutex is a mutex used to synchronize access to the client
	Log           Logger          // Log represents the logger used for logging
}

func getToken(token ...string) string {
	// If no token is provided, attempt to get the token from the config
	if len(token) == 0 {
		conf := GetConfig()
		if conf.Token == "" {
			return ""
		}
		return conf.Token
	}
	// Otherwise, use the provided token
	if len(token[0]) == 0 {
		return ""
	}
	return token[0]

}

// NewClient creates a new instance of the Client struct
func NewClient(token ...string) (*Client, error) {
	tokenStr := getToken(token...)
	if len(tokenStr) == 0 {
		return nil, FailGetToken
	}

	s := Client{}
	s.Log = GetLogger()
	s.token = tokenStr
	// Set the create time and last connection time to the current time
	s.CreateTime = time.Now().Unix()
	s.LastConnTime = time.Now().Unix()

	// Initialize the mutex and context
	s.mutex = new(sync.Mutex)
	s.Ctx = context.Background()

	// Set the logger
	s.Log = GetLogger()

	return &s, nil
}

// NewClientWithContext creates a new instance of the Client struct with the given context
func NewClientWithContext(ctx context.Context, token ...string) (*Client, error) {
	client, err := NewClient(token...)
	if err != nil {
		return nil, err
	}
	client.Ctx = ctx
	return client, nil
}

// SendMsg sends a message to the chatGPT server and returns the response
func (s *Client) SendMsg(msg string) (string, error) {
	// If the message is empty, return an empty string and no error
	if len(msg) == 0 {
		return "", nil
	}

	// Disable streaming
	s.stream = false

	// Get the request object
	req, err := s.getReq(msg)
	if err != nil {
		return "", err
	}

	// Send the request and get the response
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		s.Log.Error("http request error:%s", err.Error())
		return "", err
	}
	defer resp.Body.Close()

	// Check if the response status code indicates an error
	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusBadRequest {
		return "", s.handError(resp)
	}

	// Decode the response and append the message to the preMessage slice
	result := HttpResponse{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		s.Log.Error("Decode error:%v,originInfo:%s", err, resp.Body)
		return "", err
	}
	s.appendMessages(ChatMsg{Role: RoleAssistant, Content: result.Choices[0].Messages.Content})

	// Return the response message and no error
	return result.Choices[0].Messages.Content, nil
}
