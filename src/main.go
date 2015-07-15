package main

import (
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/sessionauth"
	"github.com/martini-contrib/sessions"
	"github.com/go-martini/martini"
	"log"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)
//TODO 코드 테스트 해보기

func main() {

	const (
		success_str = "1"
		fail_str = "0"
	)

	defer dbmap.Db.Close()

	sessionauth.RedirectUrl = "/user/session_needs"
	sessionauth.RedirectParam = ""
	addmysqldata()
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
		func(session sessions.Session, lg user_bind, req *http.Request) string {
			user := selectUser(lg.UserId)
			log.Printf("login %s %s\n", user.UserId, user.UserPw)
			if user.UserId == "" {
				return fail_str
			}
			hash_pw := hasher(lg.UserPw)
			log.Printf("hashed : %s\n", hash_pw)
			if hash_pw == user.UserPw {
				err := sessionauth.AuthenticateSession(session, &user)
				if err != nil {
					return fail_str
				}
				return success_str

			} else {
				return fail_str
			}
		})

	m.Get("/logout", sessionauth.LoginRequired, func(s sessions.Session, user sessionauth.User) string {
		sessionauth.Logout(s, user)
		return success_str
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

	//	유저정보
	m.Get("/profile/:arg", func(params martini.Params) string {
		var buf string = params["arg"]
		user_idx, _ := strconv.Atoi(buf)
		up := USER_PROFILE{}
		up.Get(int64(user_idx))
		return struct2json(up)
	})

	//메인화면 busboard

	//	m.Get("/board/buslist/:idx", sessionauth.LoginRequired,
	//		func(params martini.Params) string {
	//			var idx_str string
	//			idx_str = params["idx"]
	//			idx, err := strconv.Atoi(idx_str)
	//			if (err!=nil) {
	//				log.Print("fail to atoi")
	//				return "0"
	//			}
	//			var bus BUS
	//			arr, err := bus.list(idx)
	//			if (err!=nil) {
	//				log.Print(err)
	//				return "0"
	//			}
	//			len := len(arr)
	//			bi := make([]bus_info, len)
	//			for i := 0; i < len; i++ {
	//				bi[i].transform(arr[i])
	//			}
	//			return struct2json(bi)
	//		})
	m.Get("/board/buslist/:idx",
		func(params martini.Params) string {
			var idx_str string
			idx_str = params["idx"]
			idx, err := strconv.Atoi(idx_str)
			if (err!=nil) {
				log.Print("fail to atoi")
				return fail_str
			}
			var bus BUS
			log.Print(idx)
			//			TODO bus list수정
			arr, err := bus.list()
			if (err!=nil) {
				log.Print(err)
				return fail_str
			}
			len := len(arr)
			bi := make([]bus_info, len)
			for i := 0; i < len; i++ {
				bi[i].transform(arr[i])
			}
			return struct2json(bi)
		})
	m.Get("/board/bus/:arg",
		func(param martini.Params) string {
			var buf string = param["arg"]
			id,_ := strconv.Atoi(buf)
			var bus BUS
			bus.view(int64(id))
			if bus.Id == 0 {
				return fail_str
			}
			return struct2json(bus)
		})
	m.Post("/board/bus", binding.Bind(bus_form{}), sessionauth.LoginRequired,
		func(s sessions.Session, user sessionauth.User, bf bus_form) string {
			var bw bus_write
			u := user.(*USER_DB)
			bw.transform(bf, u.UserName, u.Id)
			bus := bw.make_bus()
			log.Print(struct2json(bus))
			bus.write()
			return success_str
		})
	m.Put("/board/bus/:arg", sessionauth.LoginRequired,
		func(s sessions.Session, user sessionauth.User, param martini.Params) string {
			var buf string
			buf = param["arg"]
			id, _ := strconv.Atoi(buf)
			var bus BUS
			bus.view(int64(id))
			if bus.Id == 0 {
				return fail_str
			}
			u := user.(*USER_DB)
			if bus.WriterId != u.Id {
				return fail_str
			}

			bus.Status = 1
			bus.update()
			return success_str
		})
	m.Get("/board/songun", func() string {
		var arr []BUS
		_, err := dbmap.Select(&arr, "select * from BUSBOARD where WriterId = 1 order by Created desc")
		check_err(err, "error")
		return struct2json(arr)
	})

	//	m.Get("/board/porfol/:arg", func(param martini.Params) string {
	//		var buf string = param["arg"]
	//		id, _ := strconv.Atoi(buf)
	//		pf := PORTFOLIO{}
	//		pf.view(int64(id))
	//		return struct2json(pf)
	//	})
	//
	//	m.Post("/board/porfol", sessionauth.LoginRequired, binding.MultipartForm(portfolio_form{}),
	//		func(s sessions.Session, user sessionauth.User, pf portfolio_form) string {
	//			image, err := pf.save_image()
	//			if err != nil {
	//				log.Print(err)
	//				return fail_str
	//			}
	//			u := user.(*USER_DB)
	//			portfolio := newPorfolio(u.Id, u.UserName, pf.Content, image)
	//			portfolio.write()
	//			return success_str
	//		})

	m.RunOnAddr(":8989")
}
