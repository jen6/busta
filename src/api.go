package main

import (
	"github.com/coopernurse/gorp"
	_"github.com/go-sql-driver/mysql"
)
//db map global uses init in util.go
var dbmap * gorp.DbMap

type user_bind struct {
	userID string `form:"id"`
	userPW string `form:"pw"`
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

func login(lg user_bind)(int, string){
	user := selectUser(lg.userID)
	if lg.userPW == user.UserPw {
		return 200, "Good!"
	} else {
		return 404, "Fuck!"
	}
}
