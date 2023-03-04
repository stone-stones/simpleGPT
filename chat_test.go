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
	mockResp = "mock_content"
	mockToken = "mock_token"
)

type MockBody struct {
	buf []byte

}
func MockOkResp() []byte {
	data := HttpResponse{}
	data.ID="test_mock_id"
	data.Object="mock_obj"
	data.Model=DefaultModelName
	data.Created=1677908733
	data.Usage.CompletionTokens=10
	data.Usage.PromptTokens=5
	data.Usage.TotalTokens=15
	data.Choices= []HttpChoice{{Index: 0,FinishReason: "stop",Messages: ChatMsg{Role: RoleAssistant,
		Content: mockResp}}}
	buf,_ := json.Marshal(data)
	return buf
}

func TestChatClient_NewClient(t *testing.T) {

	Convey("test NewClient get token fail",t,func(){

		_,err  := NewClient()
		So(err,ShouldEqual,FailGetToken)
	})

	Convey("test NewClient SendMsg",t,func(){
		patch := gomonkey.ApplyMethodFunc(http.DefaultClient,"Do",func(req *http.Request)(*http.Response, error){
			response := http.Response{}
			response.StatusCode=http.StatusOK
			response.Body =  io.NopCloser(bytes.NewReader(MockOkResp()))
			response.Body.Close()
			return &response,nil
		})
		defer patch.Reset()
		cli,err  := NewClient(mockToken)
		So(err,ShouldBeNil)
		So(cli.Token,ShouldEqual, mockToken)
		So(cli.Log,ShouldEqual,GetLogger())

		resp,err := cli.SendMsg("hello")
		So(err,ShouldBeNil)
		So(resp,ShouldEqual,mockResp)
		So(cli.Stream,ShouldEqual,false)
		So(cli.preMessage,ShouldResemble,[]ChatMsg{{RoleUser,"hello"},{RoleAssistant,mockResp}})
	})
}


func TestNewChatClient_Stream(t *testing.T) {
	cli,err  := NewClient()
	if err != nil {
		panic(err)
	}
	for _,v := range []string{"林黛玉和贾宝玉是哪个文学作品里面的","介绍一下这部作品","他们最后结局怎么样"} {
		rsp,err := cli.SendStreamMsg(v)
		if err != nil {
			t.Logf("error:%v",err)
			t.FailNow()
		}
		t.Log(rsp)
	}
}