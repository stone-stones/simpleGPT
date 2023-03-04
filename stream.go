package chatGPT

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
)


func (s *Client) SendStreamMsg(msg string)(string,error) {
	if len(msg) == 0 {
		return "",nil
	}
	s.Stream=true
	req,err := s.getReq(msg)
	if err != nil {
		return "",err
	}
	req.Header.Set("Accept", "text/event-stream")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Connection", "keep-alive")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "",err
	}
	recvCh := make(chan string,10 )
	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusBadRequest {
		return "",s.handError(resp)
	}
	go s.read(resp.Body,recvCh)
	content,err := s.recv(recvCh)
	if err != nil {
		if errors.Is(err, io.EOF) {
			return content,nil
		}
	}
	return content,err

}


func (s *Client) read(body io.ReadCloser,recvCh chan string ) {
	defer body.Close()
	var line []byte
	var err error
	var response *StreamResponse
	reader := bufio.NewReader(body)
	for {
		line,err = reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				recvCh<-StreamEndSign
				return
			}

			continue
		}
		var headerData = []byte("data: ")
		if !bytes.HasPrefix(line, headerData) {
			continue
		}
		line = bytes.TrimSuffix(line, []byte{'\n'})
		line = bytes.TrimPrefix(line, headerData)
		if  bytes.Contains(line,[]byte(StreamEndSign)){
			err = io.EOF
			recvCh<- StreamEndSign
			return
		}
		err = json.Unmarshal(line,&response)
		if err != nil {
			s.Log.Error("Unmarshal error:%s,line:%s",err,line)
			recvCh<- StreamEndSign
			return
		}
		if  len(response.Choices) >0 && len(response.Choices[0].Delta.Content)  > 0 {
			recvCh<- response.Choices[0].Delta.Content
		}
	}

}
func (s *Client) recv(recvCh chan string)(string, error)  {
	tmpMsg := ""
	for c  := range recvCh {
		if c == StreamEndSign{
			s.appendMessages(ChatMsg{Role: RoleAssistant,Content: tmpMsg})
			tmpMsg = strings.Replace(tmpMsg,"\n","",-1)
			return tmpMsg,nil
		}
		tmpMsg+=c
	}
	return  "",io.EOF
}

