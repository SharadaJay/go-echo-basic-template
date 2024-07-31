package main

import (
	"basic-echo-app/BookStore/controller"
	"basic-echo-app/BookStore/database"
	"basic-echo-app/BookStore/service"
	"basic-echo-app/BookStore/utils"
	"fmt"
	"github.com/labstack/echo"
)

func main() {

	port := utils.Cfg["server"].Port

	fmt.Println("PORT: ", port)

	echoContext := echo.New()

	datalayer := database.NewBookDatalayerImpl()

	serviceImpl := service.NewBookServiceImpl(datalayer)

	controller.NewBookController(echoContext, serviceImpl)

	echoContext.Logger.Info(echoContext.Start(port))

}
