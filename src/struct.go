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
	Want    int64        `form:"Want"`
}

type bus_write struct {
	bus_form
	Writer  string
	WriteId int64
}

func (bw *bus_write) transform(bf bus_form, name string, idx int64) {
	buf := bus_write{
		bus_form:bf,
		Writer:bus,
		WriteId:idx,
	}
	*bw = buf
}

func (bs bus_write) make_bus() BUS {
	return newBus(bs.WriteId, bs.Writer, bs.Title, bs.Content)
}

type bus_info struct {
	Id    int64
	Title string
	Name  string
}

func (bs *bus_info) transform(bus BUS) {
	bs.Title = bus.Title
	bs.Id = bus.Id
	bs.Name = bus.Writer
}