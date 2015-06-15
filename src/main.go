package main

import (
	"net/http"
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

	m.Post("/user/login", binding.Bind(user_bind{}),
		func(session sessions.Session, lg user_bind, req *http.Request) string {
			user := selectUser(lg.UserId)
			if user.UserId == "" {
				return "no User"
			}
			hash_pw := hasher(lg.UserPw)
			if hash_pw == user.UserPw {
				err := sessionauth.AuthenticateSession(session, &user)
				check_err(err, "session error")
				return "good!"

			} else {
				return "Fuck!"
			}
	})

	m.RunOnAddr(":8989")
}
