package controller

import (
	"app/src/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
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
func Index(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		db.Order("id desc").Find(&articles)
		return c.Render(http.StatusOK, "index", articles)
	}
}

//新規投稿処理
func Store(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		title := c.FormValue("title")
		content := c.FormValue("content")
		article := models.Article{Title: title, Content: content}
		db.Create(&article)
		return c.Redirect(http.StatusFound, "/")
	}
}

//編集画面表示
func Edit(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var article models.Article
		id := c.Param("id")
		db.First(&article, id)
		return c.Render(http.StatusOK, "edit", article)
	}
}

//更新処理
func Update(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var article models.Article
		id := c.Param("id")
		db.First(&article, id)
		title := c.FormValue("title")
		content := c.FormValue("content")
		article.Title = title
		article.Content = content
		db.Save(&article)
		return c.Redirect(http.StatusFound, "/")
	}
}

//削除処理
func Delete(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var article models.Article
		id := c.Param("id")
		db.First(&article, id)
		db.Delete(&article)
		return c.Redirect(http.StatusFound, "/")
	}
}
