package chatGPT

import (
	"bytes"
	"encoding/json"
	"github.com/agiledragon/gomonkey/v2"
	. "github.com/smartystreets/goconvey/convey"
	"io"
	"net/http"
	"testing"
)

var (
	mockResp  = "mock_content"
	mockToken = "mock_token"
)

type MockBody struct {
	buf []byte
}

func MockOkResp() []byte {
	data := HttpResponse{}
	data.ID = "test_mock_id"
	data.Object = "mock_obj"
	data.Model = DefaultModelName
	data.Created = 1677908733
	data.Usage.CompletionTokens = 10
	data.Usage.PromptTokens = 5
	data.Usage.TotalTokens = 15
	data.Choices = []HttpChoice{{Index: 0, FinishReason: "stop", Messages: ChatMsg{Role: RoleAssistant,
		Content: mockResp}}}
	buf, _ := json.Marshal(data)
	return buf
}

func TestChatClient_NewClient(t *testing.T) {

	Convey("test NewClient get token fail", t, func() {

		_, err := NewClient()
		So(err, ShouldEqual, FailGetToken)
	})

	Convey("test NewClient SendMsg", t, func() {
		patch := gomonkey.ApplyMethodFunc(http.DefaultClient, "Do", func(req *http.Request) (*http.Response, error) {
			response := http.Response{}
			response.StatusCode = http.StatusOK
			response.Body = io.NopCloser(bytes.NewReader(MockOkResp()))
			response.Body.Close()
			return &response, nil
		})
		defer patch.Reset()
		cli, err := NewClient(mockToken)
		So(err, ShouldBeNil)
		So(cli.token, ShouldEqual, mockToken)
		So(cli.Log, ShouldEqual, GetLogger())

		resp, err := cli.SendMsg("hello")
		So(err, ShouldBeNil)
		So(resp, ShouldEqual, mockResp)
		So(cli.stream, ShouldEqual, false)
		So(cli.preMessage, ShouldResemble, []ChatMsg{{RoleUser, "hello"}, {RoleAssistant, mockResp}})
	})
}

func TestChatClient_LoadFromConfig(t *testing.T) {
	Convey("TestChatClient_LoadFromConfig", t, func() {
		err := LoadConfigFromFile("./config.json", GetLogger())
		So(err, ShouldBeNil)
		cli, err := NewClient()
		So(err, ShouldBeNil)
		So(cli.token, ShouldEqual, "your_token")
	})

}


func TestChatClient_getMessages(t *testing.T) {
	Convey("TestChatClient_LoadFromConfig", t, func() {
		c,err  := NewClient("mock token")
		So(err,ShouldBeNil)
		c.appendMessages(&ChatMsg{Content: "this is mock content",Role: "user"})
		c.appendMessages(&ChatMsg{Content: "this is mock content",Role: "user"})
		c.appendMessages(&ChatMsg{Content: "this is mock content",Role: "user"})
		c.appendMessages(&ChatMsg{Content: "this is mock content",Role: "user"})
		c.appendMessages(&ChatMsg{Content: "this is mock content",Role: "user"})
		c.appendMessages(&ChatMsg{Content: "this is mock content",Role: "user"})
		c.getMessages(100)
		So(len(c.preMessage),ShouldEqual,4)
	})

}