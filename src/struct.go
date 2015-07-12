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

func (ub user_bind) Prepare() (string, map[string]interface{}) {
	return "SELECT * FROM USER WHERE UserId = :id AND UserPw = :pw",
	map[string]interface{}{"id": ub.UserId, "pw":ub.UserPw}
}

type user_info struct {
	Id          int64
	UserName    string
	UserSubject int
	UserGrade   int
	UserClass   int
	UserNum     int
}

func (ui user_info) make_user() USER_DB {
	return USER_DB{
		Id:ui.Id,
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
		Id:ud.Id,
		UserName:ud.UserName,
		UserSubject:ud.SUBJECT,
		UserGrade:ud.GRADE,
		UserClass:ud.CLASS,
		UserNum:ud.NUM,
	}
	*ui = it
}

type Board_find interface {
	Prepare() (string, map[string]interface{})
}

type bus_form struct {
	Title   string        `form:"Title"`
	Content string        `form:"Content"`
	Want    int64         `form:"Want"`
}

type bus_write struct {
	bus_form
	Writer  string
	WriteId int64
}

func (bw *bus_write) transform(bf bus_form, name string, idx int64) {
	buf := bus_write{
		bus_form:bf,
		Writer:name,
		WriteId:idx,
	}
	*bw = buf
}

func (bs bus_write) make_bus() BUS {
	return newBus(bs.WriteId, bs.Want, bs.Writer, bs.Title, bs.Content)
}

type bus_info struct {
	Id      int64
	Title   string
	Content string
}

func (bs *bus_info) transform(bus BUS) {
	var bus_content string
	if len(bus.Content) > 30 {
		bus_content = bus.Content[0:30]
	} else {
		bus_content = bus.Content
	}
	buf := bus_info{
		Title: bus.Title,
		Id: bus.Id,
		Content:bus_content,
	}
	*bs = buf
}

type USER_PROFILE struct {
	user_info
	PROFILE
}

func (up * USER_PROFILE) Get(UserIdx int64) {
	var (
		ud USER_DB
		ui user_info
		pf PROFILE
	)

	ud.search_one(UserIdx)
	ui.transform(ud)
	pf.Get(UserIdx)

	buf := USER_PROFILE{
		user_info:ui,
		PROFILE:pf,
	}
	*up = buf
}

