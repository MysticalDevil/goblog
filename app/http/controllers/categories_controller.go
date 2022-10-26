package controllers

import (
	"goblog/app/models/category"
	"goblog/app/requests"
	"goblog/pkg/flash"
	"goblog/pkg/route"
	"goblog/pkg/view"
	"net/http"
)

// CategoriesController 文章分类控制器
type CategoriesController struct {
	BaseController
}

// Create 文章分类创建页面
func (*CategoriesController) Create(w http.ResponseWriter, _ *http.Request) {
	view.Render(w, view.D{}, "categories.create")
}

// Store 保存文章分类
func (*CategoriesController) Store(w http.ResponseWriter, r *http.Request) {
	_category := category.Category{
		Name: r.PostFormValue("name"),
	}

	errors := requests.ValidateCategoryForm(_category)

	if len(errors) == 0 {
		_category.Create()
		if _category.ID >0 {
			flash.Success("分类创建成功")
			indexURL := route.Name2URL("home")
			http.Redirect(w, r, indexURL, http.StatusFound)
		} else {
			flash.Danger("创建文章分类失败，请联系管理员")
			indexURL := route.Name2URL("home")
			http.Redirect(w, r, indexURL, http.StatusInternalServerError)
		}
	} else {
		view.Render(w, view.D{
			"Category": _category,
			"Errors": errors,
		}, "categories.create")
	}
}
