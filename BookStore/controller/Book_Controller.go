package controller

import (
	"basic-echo-app/BookStore/intf"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type BookController struct {
	bookService intf.BookService
}

func NewBookController(echo *echo.Echo, bookServiceObject intf.BookService) {

	bookControllerObject := &BookController{
		bookService: bookServiceObject,
	}

	echo.GET("/printAuthor", bookControllerObject.PrintAuthor)
	echo.GET("/test", bookControllerObject.Test)
}

func (controllerObj *BookController) PrintAuthor(ec echo.Context) error {

	return ec.String(http.StatusOK, "Author X")
}

func (controllerObj *BookController) Test(ec echo.Context) error {

	fmt.Println("**** Inside Book Controller ****")

	requestContext := ec.Request().Context()
	controllerObj.bookService.TestBookService(requestContext)

	return ec.String(http.StatusOK, "Test Successful")
}
