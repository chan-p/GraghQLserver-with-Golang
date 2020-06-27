package main

import (
	"log"
	_ "net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	err := e.Start(":3000")
	if err != nil {
		log.Fatalln(err)
	}
}