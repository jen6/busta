package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/sessionauth"
	"github.com/martini-contrib/sessions"
	"github.com/martini-contrib/render"
	"github.com/coopernurse/gorp"
	_ "github.com/go-sql-driver/mysql"
)


//martini and db map global uses init in util.go
var (
	m martini.ClassicMartini = martini.Classic()
	dbmap *gorp.DbMap
)




type user_bind struct {
	UserId     string `form:"Id"`
	UserPw     string `form:"Pw"`
	unexported string `form:"-"` // skip binding of unexported fields
}

func martini_setup(m *martini.ClassicMartini, store_buf *sessions.CookieStore) {
	//TODO: m use 정리하기
	m.Use(render.Renderer())
	m.Use(sessions.Sessions("my_session", store_buf))
	m.Use(sessionauth.SessionUser(GenerateAnonymousUser))
}

func GenerateAnonymousUser() sessionauth.User {
	return &USER_DB{}
}
