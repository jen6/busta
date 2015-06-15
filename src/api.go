package main

import (
	"log"
	"github.com/codegangsta/martini-contrib/sessionauth"
	"github.com/coopernurse/gorp"
	_ "github.com/go-sql-driver/mysql"
)

//db map global uses init in util.go
var dbmap *gorp.DbMap

type user_bind struct {
	UserId     string `form:"Id"`
	UserPw     string `form:"Pw"`
	unexported string `form:"-"` // skip binding of unexported fields
}

//
//func (bp user_bind) Validate(errors *binding.Errors, req *http.Request) {
//	if req.Header.Get("X-Custom-Thing") == "" {
//		errors.Overall["x-custom-thing"] = "The X-Custom-Thing header is required"
//	}
//	if len(bp.userID) < 4 {
//		errors.Fields["title"] = "Too short; minimum 4 characters"
//	} else if len(bp.userID) > 20 {
//		errors.Fields["title"] = "Too long; maximum 20 characters"
//	}
//	if len(bp.userID) < 8 {
//		errors.Fields["title"] = "Too short; minimum 8 characters"
//	} else if len(bp.userID) > 20 {
//		errors.Fields["title"] = "Too long; maximum 20 characters"
//	}
//}

func login(lg user_bind) string {
	log.Println("login handler")
	user := selectUser(lg.UserId)
	log.Printf("%s %s", lg.UserId, lg.UserPw)
	log.Printf("%s %s", user.UserId, user.UserPw)
	if user.UserId == "" {
		return "no User"
	}
	hash_pw := hasher(lg.UserPw)
	if hash_pw == user.UserPw {
		return "Good!"
	} else {
		return "Fuck!"
	}
}

func GenerateAnonymousUser() sessionauth.User {
	return &USER_DB{}
}