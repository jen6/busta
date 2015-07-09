package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/sessionauth"
	"github.com/martini-contrib/sessions"
	"github.com/martini-contrib/render"
	"gopkg.in/gorp.v1"
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

func GenerateAnonymousUser() sessionauth.User {
	return &USER_DB{}
}
