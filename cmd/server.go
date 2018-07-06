package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"vote"
	"vote/utils"
	"net/http"
	"os"
)

func main() {
	dataPath := os.Getenv("DATA_PATH")
	file, err := os.OpenFile(dataPath+"main.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err == nil {
		utils.Logger.Out = file
	} else {
		utils.Logger.Fatalln("Failed to log to file, using default stderr")
	}

	accessLogFile, err := os.OpenFile(dataPath+"access.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		utils.Logger.Fatalln("Failed to log to file, using default stderr")
	}

	e := echo.New()
	ctl := web.NewControl(e)
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: accessLogFile,
	}))
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://vote.gmzhang.com"},
		AllowMethods:     []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
		AllowCredentials: true,
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})

	activity := e.Group("/api")
	ctl.Dispatch(activity)

	e.Logger.Fatal(e.Start(":80"))
}
