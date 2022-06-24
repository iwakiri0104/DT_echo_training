package controller

import (
	"app/src/models"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"html/template"
	"io"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func Init() *echo.Echo {

	list, err := template.New("t").ParseGlob("views/template/*.html")

	t := &Template{
		templates: template.Must(list, err),
	}

	e := echo.New()

	//メソッド対応処理
	e.Pre(MethodOverride)
	e.Renderer = t

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	//自動マイグレーション
	models.Init()
	
	db := models.DatabaseConnection()
	//router
	e.GET("/", Index(db))
	e.POST("/store", Store(db))
	e.GET("/edit/:id", Edit(db))
	e.PUT("/update/:id", Update(db))
	e.DELETE("/delete/:id", Delete(db))

	return e

}
