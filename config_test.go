package chatGPT

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestLoadConfigFromFile(t *testing.T) {
	Convey("test get config ok", t, func() {
		err := LoadConfigFromFile("./config.json", GetLogger())
		So(err, ShouldBeNil)
		config := GetConfig()
		So(config.Token, ShouldEqual, "your_token")
		So(config.PostParams.User, ShouldEqual, "simpleGPT")
	})
	Convey("test config unchanged", t, func() {
		err := LoadConfigFromFile("./config.json", GetLogger())
		So(err, ShouldBeNil)
		config := GetConfig()
		So(config.Token, ShouldEqual, "your_token")
		So(config.PostParams.User, ShouldEqual, "simpleGPT")
		config.Token = "other_token"
		So(config.Token, ShouldEqual, "other_token")
		config = GetConfig()
		So(config.Token, ShouldEqual, "your_token")

	})
}
