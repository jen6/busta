package main
import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type User_Interface interface {
	make_user() USER_DB
	Prepare() *sql.Stmt
}

type user_bind struct {
	UserId     string `form:"Id"`
	UserPw     string `form:"Pw"`
	unexported string `form:"-"` // skip binding of unexported fields
}

func (ub * user_bind) make_user() USER_DB {
	return USER_DB{
		UserId: ub.UserId,
		UserPw: ub.UserPw,
	}
}

//TODO map interface를 이용해 어떻게 쿼리를 날릴껀지 생각 해 보기
func (ub * user_bind) Prepare() (string, map[string]interface{}) {
	return ("", )
}

type user_info struct {
	UserName    string
	UserSubject string
	UserGrade   int64
	UserClass   int64
	UserNum     int64
}

func (ui * user_info) make_user() USER_DB {
	return USER_DB{
		UserName:ui.UserName,
		SUBJECT:ui.UserSubject,
		CLASS:ui.UserClass,
		GRADE:ui.UserGrade,
		NUM:ui.UserNum,
	}
}
