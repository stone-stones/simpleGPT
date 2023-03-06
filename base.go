package chatGPT

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// getPostParams returns the JSON-encoded post parameters
func (s *Client) getPostParams() ([]byte, error) {
	// Get the post parameters from the configuration
	params := GetConfig().PostParams
	if s.stream {
		params.Stream = true
	}

	params.Message = s.getMessages(params.MaxTokens)

	// Set the User parameter if it's not empty
	if s.User != "" {
		params.User = s.User
	}

	// Marshal the parameters into a byte slice
	c, err := json.Marshal(params)
	return c, err
}

// CutPreMessages cuts down the number of preMessages to remain
func (s *Client) CutPreMessages(remain int) {
	if len(s.preMessage) <= remain {
		return
	}

	// Lock the preMessage list and cut it down
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.preMessage = s.preMessage[len(s.preMessage)-remain:]
}

// getReq returns an http.Request object to send to the server
func (s *Client) getReq(msg string) (*http.Request, error) {
	// Append the user's message to the preMessage list
	s.appendMessages(&ChatMsg{Role: RoleUser, Content: msg})

	// Get the post parameters and return an http.Request object
	c, err := s.getPostParams()
	if err != nil {
		s.Log.Error("getPostParams error:%v", err)
		return nil, err
	}
	req, err := http.NewRequestWithContext(s.Ctx, http.MethodPost, GPTURL, bytes.NewReader(c))
	if err != nil {
		s.Log.Error("NewRequestWithContext error:%v", err)
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+s.token)
	return req, nil
}

// appendMessages appends one or more messages to the preMessage list
func (s *Client) appendMessages(msg ...*ChatMsg) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.preMessage = append(s.preMessage, msg...)
}

// getMessages gets the messages for params
func (s *Client) getMessages(maxToken int) []*ChatMsg {
	// Get the preMessage list and lock it to prevent race conditions
	s.mutex.Lock()
	defer s.mutex.Unlock()

	var totalRune int
	var messages []*ChatMsg

	for i := len(s.preMessage) - 1; i >= 0; i-- {
		msg := s.preMessage[i]
		runeCount := len([]rune(msg.Content)) + len([]rune(msg.Role))
		if totalRune+runeCount > maxToken {
			break
		}
		totalRune += runeCount
		messages = append([]*ChatMsg{msg}, messages...)
	}
	s.preMessage = messages
	return messages
}


// handError handles an http response with an error status code
func (s *Client) handError(resp *http.Response) error {
	// Decode the error response
	var errRes ErrorResponse
	err := json.NewDecoder(resp.Body).Decode(&errRes)

	// If there's an error decoding the response or the error message is empty,
	// create a RequestError and return it
	if err != nil || errRes.Error == nil {
		reqErr := RequestError{
			StatusCode: resp.StatusCode,
			Err:        err,
		}
		return fmt.Errorf("error, %v", &reqErr)
	}

	// Add the status code to the error message and return it
	errRes.Error.StatusCode = resp.StatusCode
	s.Log.Error("error, status code: %d, message: %s", resp.StatusCode, errRes.Error)
	return errRes.Error
}
