package conncheck

import (
	"testing"

	c "github.com/smartystreets/goconvey/convey"
)

func TestCheckUrl(t *testing.T) {
	c.Convey("基础用例", t, func() {
		// 测试
		url := "https://www.baidu.com/"
		got := CheckUrl(url)

		// 断言
		c.So(got, c.ShouldEqual, true) // 两种方式都可以
		c.ShouldBeTrue(got)
	})

	c.Convey("url请求不通的示例", t, func() {
		url := "/test/ssss"
		got := CheckUrl(url)

		c.ShouldBeFalse(got)
	})
}
