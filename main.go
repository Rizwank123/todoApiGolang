package main

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/rizwank123/controller"
	"github.com/rizwank123/database"
	_ "github.com/rizwank123/docs"
	"github.com/rizwank123/repository"
	"github.com/rizwank123/route"
	"github.com/rizwank123/service"
)

// @title						ToDo Api
// @version					1.0
// @description				go lang practice rest api.
// @termsOfService				http://swagger.io/terms/
// @contact.name				API Support
// @contact.url				http://www.swagger.io/support
// @contact.email				support@swagger.io
// @license.name				Apache 2.0
// @license.url				http://www.apache.org/licenses/LICENSE-2.0.html
// @host						localhost:8090
// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
// @BasePath
func main() {
	e := echo.New()

	dbs := database.DatabseInit()

	//user rpo initialize
	ur := repository.NewUserRepository(dbs)
	//user service intialize
	us := service.NewService(ur)
	//user controller initialize
	uc := controller.NewUserController(us)
	// task Repository initialize
	tskRepo := repository.NewTaskRepo(dbs)
	taskSer := service.NewTaskService(tskRepo, ur)
	tskCont := controller.NewTaskController(taskSer)

	//api initialize
	api := route.NewApi(uc, tskCont)
	//route intilaize
	api.IntitRoute(e)

	// swagger-ui
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	//init swagger

	defer dbs.Close()

	e.Logger.Fatal(e.Start("localhost:8090"))
}
