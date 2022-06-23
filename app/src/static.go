package main

import "github.com/labstack/echo/v4"

// 静的ファイルを配置するルーティングを設定
func SetStaticRoute(e *echo.Echo) {
	e.Static("/css", "./views/css")
	e.Static("/js", "./views/js")
}
