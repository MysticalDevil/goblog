package controllers

import (
	"errors"
	"fmt"
	"goblog/pkg/auth"
	"goblog/pkg/flash"
	"goblog/pkg/logger"
	"goblog/pkg/route"
	"goblog/pkg/view"
	"gorm.io/gorm"
	"net/http"
	"strconv"

	"goblog/app/models/user"
	"goblog/app/requests"
)

// AuthController 处理用户认证
type AuthController struct {
}

// Register 注册页面
func (*AuthController) Register(w http.ResponseWriter, _ *http.Request) {
	view.RenderSimple(w, view.D{}, "auth.register")
}

// DoRegister 处理注册逻辑
func (*AuthController) DoRegister(w http.ResponseWriter, r *http.Request) {
	_user := user.User{
		Name:            r.PostFormValue("name"),
		Email:           r.PostFormValue("email"),
		Password:        r.PostFormValue("password"),
		PasswordConfirm: r.PostFormValue("password_confirm"),
	}

	errs := requests.ValidateRegistrationForm(_user)

	if len(errs) > 0 {
		view.RenderSimple(w, view.D{
			"Errors": errs,
			"User":   _user,
		}, "auth.register")
	} else {
		err := _user.Create()
		if err != nil {
			logger.LogError(err)
		}

		if _user.ID > 0 {
			flash.Success("恭喜您注册成功！")
			auth.Login(_user)
			http.Redirect(w, r, "/", http.StatusFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "注册失败，请联系管理员")
		}
	}
}

// Login 登录页面
func (*AuthController) Login(w http.ResponseWriter, _ *http.Request) {
	view.RenderSimple(w, view.D{}, "auth.login")
}

// DoLogin 处理登录逻辑
func (*AuthController) DoLogin(w http.ResponseWriter, r *http.Request) {
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")

	if err := auth.Attempt(email, password); err == nil {
		flash.Success("欢迎回来！")
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		view.RenderSimple(w, view.D{
			"Error":    err.Error(),
			"Email":    email,
			"Password": password,
		}, "auth.login")
	}
}

// Logout 用户登出
func (*AuthController) Logout(w http.ResponseWriter, r *http.Request) {
	auth.Logout()
	flash.Success("您已退出登录！")
	http.Redirect(w, r, "/", http.StatusFound)
}

// SendEmail 发送邮件表单
func (*AuthController) SendEmail(w http.ResponseWriter, _ *http.Request) {
	view.RenderSimple(w, view.D{}, "auth.sendEmail")
}

// DoSendEmail 发送邮件
func (*AuthController) DoSendEmail(w http.ResponseWriter, r *http.Request) {
	email := r.PostFormValue("email")

	// 根据 Email 获取用户
	_user, err := user.GetByEmail(email)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = errors.New("账号不存在")
		} else {
			err = errors.New("系统内部错误，请稍后重试")
		}
		view.RenderSimple(w, view.D{
			"Error": err.Error(),
			"Email": email,
		}, "auth.sendEmail")
	} else {
		id := _user.ID
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "邮件发送成功，ID: " + strconv.FormatUint(id, 10))
	}
}

// ResetPassword 重置密码表单
func (*AuthController) ResetPassword(w http.ResponseWriter, r *http.Request) {
	id := route.GetRouteVariable("id", r)
	view.RenderSimple(w, view.D{
		"ID": id,
	}, "auth.reset")
}

// DoResetPassword 重置密码
func (*AuthController) DoResetPassword(w http.ResponseWriter, r *http.Request) {
	id := r.PostFormValue("id")
	_user, err := user.Get(id)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "404 用户未找到")
		} else {
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 服务器内部错误")
		}
	} else {
		password := r.PostFormValue("password")
		passwordConfirm := r.PostFormValue("password_confirm")

		errs := validateResetPasswordFormData(password, passwordConfirm)
		if len(errs) > 0 {
			view.RenderSimple(w, view.D{
				"ID": id,
				"Errors": errs,
				"Password": password,
				"PasswordConfirm": passwordConfirm,
			}, "auth.reset")
		} else {
			_user.Password = password
			rowsAffected, err := _user.Update()
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprint(w, "500 服务器内部错误")
				return
			}
			if rowsAffected >0  {
				showURL := route.Name2URL("auth.login")
				http.Redirect(w, r, showURL, http.StatusFound)
			} else {
				fmt.Fprint(w, "您没有做任何更改！")
			}
		}
	}
}

func validateResetPasswordFormData(password, passwordConfirm string) map[string][]string {
	errs := make(map[string][]string)
	if password == "" {
		errs["password"] = append(errs["password"], "密码为必填项")
	}
	if len(password) < 6 {
		errs["password"] = append(errs["password"], "密码长度需大于 6")
	}
	if password != passwordConfirm {
		errs["password_confirm"] = append(errs["password_confirm"], "两次输入密码不匹配！")
	}
	return errs
}
