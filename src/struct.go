package main
import (

)

type User_Interface interface {
	make_user() USER_DB
	Prepare() (string, map[string]interface{})
}



type user_bind struct {
	UserId     string `form:"Id"`
	UserPw     string `form:"Pw"`
	unexported string `form:"-"` // skip binding of unexported fields
}

func (ub user_bind) make_user() USER_DB {
	return USER_DB{
		UserId: ub.UserId,
		UserPw: ub.UserPw,
	}
}

//TODO map interface를 이용해 어떻게 쿼리를 날릴껀지 생각 해 보기
func (ub user_bind) Prepare() (string, map[string]interface{}) {
	return "SELECT * FROM USER WHERE UserId = :id AND UserPw = :pw",
	map[string]interface{}{"id": ub.UserId, "pw":ub.UserPw}
}

type user_info struct {
	UserName    string
	UserSubject int
	UserGrade   int
	UserClass   int
	UserNum     int
}

func (ui user_info) make_user() USER_DB {
	return USER_DB{
		UserName:ui.UserName,
		SUBJECT:ui.UserSubject,
		CLASS:ui.UserClass,
		GRADE:ui.UserGrade,
		NUM:ui.UserNum,
	}
}

func (ui user_info) Prepare() (string, map[string]interface{}) {
	return "SELECT * FROM USER WHERE UserName = :name", map[string]interface{}{"name" : ui.UserName}
}

func (ui * user_info) transform(ud USER_DB) {
	it := user_info{
		UserName:ud.UserName,
		UserSubject:ud.SUBJECT,
		UserGrade:ud.GRADE,
		UserClass:ud.CLASS,
		UserNum:ud.NUM,
	}
	ui = &it
}

