package main

import (
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/sessionauth"
	"github.com/martini-contrib/sessions"
	"github.com/go-martini/martini"
	"net/http"
	"log"
	"github.com/martini-contrib/render"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	m := martini.Classic()
	defer dbmap.Db.Close()

	store := sessions.NewCookieStore([]byte("secret123"))
	store.Options(sessions.Options{
		MaxAge: 0,
	})
	//TODO: m use 정리하기
	m.Use(render.Renderer())
	m.Use(sessions.Sessions("my_session", store))
	m.Use(sessionauth.SessionUser(GenerateAnonymousUser))
	sessionauth.RedirectUrl = "/user/login"
	sessionauth.RedirectParam = "new-next"

	m.Get("/json_test", func() string {
		return "hello"
	})

	//	TODO:TEST확인
	m.Get("/tester", sessionauth.LoginRequired, func(user sessionauth.User) string {
		a := user.(*USER_DB)
		return a.UserName
	})

	m.Post("/user/login", binding.Bind(user_bind{}),
		func(session sessions.Session, lg user_bind, r render.Render, req *http.Request) {
			user := selectUser(lg.UserId)
			log.Printf("login %s %s\n", user.UserId, user.UserPw)
			if user.UserId == "" {
				r.Redirect(sessionauth.RedirectUrl)
				return 
			}
			hash_pw := hasher(lg.UserPw)
			log.Printf("hashed : %s\n", hash_pw)
			if hash_pw == user.UserPw {
				err := sessionauth.AuthenticateSession(session, &user)
				if err != nil{
					return 
				}
				params := req.URL.Query()
				redirect := params.Get(sessionauth.RedirectParam)
				r.Redirect(redirect)
				return 

			} else {
				return 
			}
		})

		m.Get("/user/logout", sessionauth.LoginRequired, func(s sessions.Session, user sessionauth.User) string {
		sessionauth.Logout(s, user)
		return "logout"
	})

	m.RunOnAddr(":8989")
}
