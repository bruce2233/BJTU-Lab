package decorator

import "time"

type UserImpl struct {
}

type LoginExtendedImpl struct {
	IUser
}

type Teacher struct {
	IUser
}

type Student struct {
	IUser
}

type Graduate struct {
	IUser
}

type IUser interface {
	login()
	logout()
}

func (user UserImpl) login() {
	println("login")
}

func (user UserImpl) logout() {
	println("logout")
}

func (userExtend LoginExtendedImpl) login() {
	println("日志记录登录时间:", time.Now().Format("2006-01-02 15:04:05"))
	print("调用User:")
	userExtend.IUser.login()
}

func (userExtend LoginExtendedImpl) logout() {
	println("日志记录登出时间:", time.Now().Format("2006-01-02 15:04:05"))
	print("调用User:")
	userExtend.IUser.logout()
}
