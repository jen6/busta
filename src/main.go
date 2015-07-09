package main

import (
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/sessionauth"
	"github.com/martini-contrib/sessions"
	"github.com/martini-contrib/render"
	"github.com/go-martini/martini"
	"log"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)
//TODO 코드 테스트 해보기

func main() {

	defer dbmap.Db.Close()

	sessionauth.RedirectUrl = "/user/session_needs"
	sessionauth.RedirectParam = ""

	m.Get("/user/session_needs", func() string {
		return "NULL"
	})

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
				if err != nil {
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
	//TODO 나중에 테스트 끝나면 유저 찾는부분에 세션 인증하는 부분 넣기
	//TODO 유저 프로필 같이 가져오는 기능 만들기
	m.Get("/user/:name", func(params martini.Params) string {
		var name string = params["name"]
		log.Print(name)
		user := user_info{
			UserName: name,
		}
		var user_search USER_DB
		user_arr := user_search.search_arr(user)
		len := len(user_arr)
		info_arr := make([]user_info, len)
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

		var ui user_info
		ui.transform(user_search)

		return struct2json(ui)
	})

	//메인화면 busboard
	//TODO post글쓰기 테스트 하기
	m.Post("/board/bus", binding.Bind(bus_form{}), sessionauth.LoginRequired,
		func(s sessions.Session, user sessionauth.User, bf bus_form) string {
			var bw bus_write
			u := user.(*USER_DB)
			bw.transform(bf, u.UserName, u.Id)
			return "Ok"
		})
	m.RunOnAddr(":8989")
}
