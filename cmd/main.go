package main

import (
	"fmt"
	"github.com/kataras/iris/v12"
	irisLog "github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/kataras/iris/v12/mvc"
	"gitlab.feedtoken.tech/xdance_group/xdance/config"
	"gitlab.feedtoken.tech/xdance_group/xdance/internal/controller"
	"gitlab.feedtoken.tech/xdance_group/xdance/internal/repository"
	"gitlab.feedtoken.tech/xdance_group/xdance/internal/service"
	"gitlab.feedtoken.tech/xdance_group/xdance/pkg/database"
	"gitlab.feedtoken.tech/xdance_group/xdance/pkg/logger"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	//load config
	config.SetUp()
	//load logs
	logger.SetUp()
	//load db
	db = database.SetUp()

}

func main() {
	app := iris.New()
	app.Use(recover.New())
	app.Use(irisLog.New())

	app.Party("/api/v1").ConfigureContainer(container)
	mvc.Configure(app.Party("/api/v1"), mvcConfig)

	port := config.Ops.Server.Port
	app.Listen(fmt.Sprintf(":%d", port), iris.WithOptimizations)
}

func mvcConfig(app *mvc.Application) {
	//添加依赖
	app.Register(service.NewSignService(new(repository.User)))
	app.Handle(new(controller.SignController))
}

func container(api *iris.APIContainer) {
	api.RegisterDependency(db)
	api.Post("/sign/up", controller.SignUp)
}
