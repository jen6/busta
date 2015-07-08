package main

import (
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/sessionauth"
	"github.com/martini-contrib/sessions"
	"github.com/martini-contrib/render"
	"github.com/go-martini/martini"
	"net/http"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)
//TODO 코드 테스트 해보기
//TODO 데이터 구조 바꾸기 DB에 사용되는 구조체들에 대해 INTERFACE 빼주고 동일한 동작(SEARCH)같은 메서드 정의

func main() {

	defer dbmap.Db.Close()

	sessionauth.RedirectUrl = "/user/login"
	sessionauth.RedirectParam = "new-next"

	m.Get("/json_test", func() string {
		return "hello"
	})

	m.Get("/tester", func(user sessionauth.User) string {
		a := user.(*USER_DB)
		return a.UserName
	})

	m.Post("/login", binding.Bind(user_bind{}),
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

	m.Get("/logout", sessionauth.LoginRequired, func(s sessions.Session, user sessionauth.User) string {
		sessionauth.Logout(s, user)
		return "logout"
	})

	m.Get("/user/:name", func(params martini.Params) string {
		var name string = params["name"]
		log.Print(name)
		user := user_info{
			UserName: name,
		}
		var user_search USER_DB
		user_arr := user_search.search_arr(user)
		len := len(user_arr)
		var info_arr [len]user_info
		for i := 0; i < len; i++ {
			info_arr[i].transform(user_arr[i])
		}
		return struct2json(info_arr)
	})

	m.Post("/user/:idx", func(params martini.Params) string {
		var buf string = params["idx"]
		idx, err := strconv.Atoi(buf)
		if (err!=nil) {
			log.Print("fail to atoi")
			return "NULL"
		}
		var user_search USER_DB
		user_search.GetById(idx)
		return struct2json(user_search)
	})

	m.RunOnAddr(":8989")
}
