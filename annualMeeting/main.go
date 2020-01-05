package annualMeeting

import (
	"github.com/golang/gofrontend/libgo/go/fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

var userList []string

type lotteryController struct {
	Ctx iris.Context
}

func newApp() *iris.Application {
	app := iris.New()
	mvc.New(app.Party("/")).Handle(&lotteryController{})

	return app
}

func main() {
	app := newApp()
	userList = []string{}

	// 启动 8080 端口
	app.Run(iris.Addr(":8080"))
}

func (c *lotteryController) Get() string {
	count := len(userList)

	return fmt.Sprint("参与抽奖的总用户数: %d\n", count)
}

func 






