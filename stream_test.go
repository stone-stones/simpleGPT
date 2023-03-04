package chatGPT

import (
	"bytes"
	"fmt"
	"github.com/agiledragon/gomonkey/v2"
	. "github.com/smartystreets/goconvey/convey"
	"io"
	"net/http"
	"strings"
	"testing"
	"encoding/json"
)


var mockRes  = "Hello! How can I assist you today??"
func MockStreamOkResp() []byte {

	buf := bytes.NewBufferString("")
	tmp := `data: {"id":"chatcmpl","object":"chat.completion.chunk","created":1677920385,"model":"gpt-3.5-turbo-0301","choices":[{"delta":{"content":" %s"},"index":0,"finish_reason":null}]}
`
	for _,v := range strings.Split(mockRes+" "+StreamEndSign," ") {
		buf.WriteString(fmt.Sprintf(tmp,v))
	}
	return buf.Bytes()
}




func MockFailResp() []byte {
	data := APIError{}
	data.Message = "mock error message"
	data.Type="invalid_request_error"
	code := "invalid_api_key"
	data.Code=&code
	buf,_ :=json.Marshal(data)
	return buf
}

func TestChatClient_SendStreamMsg(t *testing.T) {
	Convey("test NewClient SendStreamMsg ok",t,func(){
		patch := gomonkey.ApplyMethodFunc(http.DefaultClient,"Do",	 func(req *http.Request)(*http.Response, error){
			response := http.Response{}
			response.StatusCode=http.StatusOK
			response.Body =io.NopCloser(bytes.NewReader(MockStreamOkResp()))
			return &response,nil
		})
		defer patch.Reset()
		cli,err  := NewClient(mockToken)
		So(err,ShouldBeNil)
		So(cli.Token,ShouldEqual, mockToken)
		So(cli.Log,ShouldEqual,GetLogger())

		resp,err := cli.SendStreamMsg("hello")
		So(err,ShouldBeNil)
		So(resp,ShouldContainSubstring,mockRes)
		So(cli.Stream,ShouldEqual,true)
		So(cli.preMessage,ShouldResemble,[]ChatMsg{{RoleUser,"hello"},{RoleAssistant," "+mockRes}})
	})
	Convey("test NewClient SendStreamMsg fail",t,func(){
		patch := gomonkey.ApplyMethodFunc(http.DefaultClient,"Do",func(req *http.Request)(*http.Response, error){
			response := http.Response{}
			response.StatusCode=http.StatusBadRequest
			response.Body =io.NopCloser(bytes.NewReader(MockFailResp()))
			return &response,nil
		})
		defer patch.Reset()
		cli,err  := NewClient(mockToken)
		So(err,ShouldBeNil)
		So(cli.Token,ShouldEqual, mockToken)
		So(cli.Log,ShouldEqual,GetLogger())

		_,err = cli.SendStreamMsg("hello")
		So(err,ShouldNotBeNil)

	})

}
