package main

import (
	"github.com/codegangsta/martini-contrib/binding"
	"github.com/go-martini/martini"
	"github.com/codegangsta/martini-contrib/sessionauth"
	"github.com/codegangsta/martini-contrib/sessions"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	m := martini.Classic()
	defer dbmap.Db.Close()

	store := sessions.NewCookieStore([]byte("secret123"))
	store.Options(sessions.Options{
		MaxAge: 0,
	})
	m.Use(sessions.Sessions("my_session", store))
	m.Use(sessionauth.SessionUser(GenerateAnonymousUser))
	sessionauth.RedirectUrl = "/new-login"
	sessionauth.RedirectParam = "new-next"



	m.Get("/json_test", func() string {
		return "hello"
	})

	m.Post("/user/login", binding.Bind(user_bind{}), func(lg user_bind) string {
		return login(lg)
	})

	m.RunOnAddr(":8989")
}
