package chatGPT

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *Client)getPostParams() ([]byte,error) {
	params := GetConfig().PostParams
	if s.Stream {
		params.Stream=true
	}
	s.mutex.Lock()
	params.Message = s.preMessage
	s.mutex.Unlock()
	if s.User != "" {
		params.User = s.User
	}
	c ,err := json.Marshal(params)
	return c ,err

}

func (s *Client)CutPreMessages(remain int) {
	if len(s.preMessage) <= remain {
		return
	}
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.preMessage = s.preMessage[len(s.preMessage)-remain:]
}

func (s *Client)getReq(msg string )(*http.Request,error ) {
	s.appendMessages(ChatMsg{Role: RoleUser,Content: msg})
	c,err := s.getPostParams()
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(s.Ctx,http.MethodPost, GPTURL, bytes.NewReader(c))
	if err != nil {

		return nil,err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+s.Token)
	return req,nil
}

func (s *Client) appendMessages(msg ...ChatMsg) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.preMessage = append(s.preMessage,msg...)
}

func (s *Client) handError(resp *http.Response) error  {
	var errRes ErrorResponse
	err := json.NewDecoder(resp.Body).Decode(&errRes)
	if err != nil || errRes.Error == nil {
		reqErr := RequestError{
			StatusCode: resp.StatusCode,
			Err:        err,
		}
		return fmt.Errorf("error, %v", &reqErr)
	}
	errRes.Error.StatusCode = resp.StatusCode
	s.Log.Error("error, status code: %d, message: %s", resp.StatusCode, errRes.Error)
	return errRes.Error
}