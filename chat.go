package chatGPT

import (
	"context"
	"encoding/json"
	"net/http"
	"sync"
	"time"
)

type Client struct {
	User          string
	Token         string
	Stream        bool
	MaxPreMessage int
	CreateTime    int64
	LastConnTime  int64
	preMessage    []ChatMsg
	Ctx           context.Context
	mutex         *sync.Mutex
	Log           Logger
}


func NewClient(token ...string)(*Client,error) {
	s := Client{}
	s.Log=GetLogger()
	if len(token) == 0 {
		conf := GetConfig()
		if conf.Token == "" {
			return nil,FailGetToken
		}
		s.Token=conf.Token
	} else {
		if len(token[0]) == 0 {
			return nil,FailGetToken
		}
		s.Token=token[0]
	}
	s.CreateTime = time.Now().Unix()
	s.LastConnTime = time.Now().Unix()
	s.mutex=new(sync.Mutex)
	s.Ctx = context.Background()
	s.Log=GetLogger()
	return &s ,nil
}

func NewClientWithContext(ctx context.Context,token ...string)(*Client,error) {
	client,err := NewClient(token...)
	if err != nil {
		return nil,err
	}
	client.Ctx = ctx
	return client,nil
 }


func (s *Client) SendMsg(msg string)(string, error) {
	if len(msg) == 0 {
		return "",nil
	}
	s.Stream=false
	req,err := s.getReq(msg)
	if err != nil {
		return "",err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		s.Log.Error("http request error:%s",err.Error())
		return "",err
	}
	defer resp.Body.Close()

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusBadRequest {
		return "",s.handError(resp)
	}
	result :=HttpResponse{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err !=nil {
		s.Log.Error("Decode error:%v",err)
		return "",err
 	}
	s.appendMessages(ChatMsg{Role: RoleAssistant,Content: result.Choices[0].Messages.Content})
	return result.Choices[0].Messages.Content,nil

}

