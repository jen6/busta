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
	m *martini.ClassicMartini = martini.Classic()
	dbmap *gorp.DbMap
)

func init() {
	store := sessions.NewCookieStore([]byte("secret123"))
	store.Options(sessions.Options{
		MaxAge: 0,
	})
	m.Use(render.Renderer())
	m.Use(sessions.Sessions("my_session", store))
	m.Use(sessionauth.SessionUser(GenerateAnonymousUser))
}



type user_bind struct {
	UserId     string `form:"Id"`
	UserPw     string `form:"Pw"`
	unexported string `form:"-"` // skip binding of unexported fields
}


func GenerateAnonymousUser() sessionauth.User {
	return &USER_DB{}
}
