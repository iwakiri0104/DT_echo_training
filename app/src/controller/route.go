package controller

import (
	"app/src/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

var articles []models.Article

//メソッド対応処理
func MethodOverride(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Request().Method == "POST" {
			method := c.Request().PostFormValue("_method")
			if method == "PUT" || method == "PATCH" || method == "DELETE" {
				c.Request().Method = method
			}
		}
		return next(c)
	}
}

//一覧表示
func Index(c echo.Context) error {
	db := models.DatabaseConnection()
	db.Order("id desc").Find(&articles)
	return c.Render(http.StatusOK, "index", articles)
}

//新規投稿処理
func Store(c echo.Context) error {
	title := c.FormValue("title")
	content := c.FormValue("content")
	db := models.DatabaseConnection()
	article := models.Article{Title: title, Content: content}
	db.Create(&article)
	return c.Redirect(http.StatusFound, "/")
}

//編集画面表示
func Edit(c echo.Context) error {
	var article models.Article
	id := c.Param("id")
	db := models.DatabaseConnection()
	db.First(&article, id)
	return c.Render(http.StatusOK, "edit", article)
}

//更新処理
func Update(c echo.Context) error {
	var article models.Article
	db := models.DatabaseConnection()
	id := c.Param("id")
	db.First(&article, id)
	title := c.FormValue("title")
	content := c.FormValue("content")
	article.Title = title
	article.Content = content
	db.Save(&article)
	return c.Redirect(http.StatusFound, "/")
}

//削除処理
func Delete(c echo.Context) error {
	var article models.Article
	db := models.DatabaseConnection()
	id := c.Param("id")
	db.First(&article, id)
	db.Delete(&article)
	return c.Redirect(http.StatusFound, "/")
}
