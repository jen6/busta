package main

import (
	"github.com/go-martini/martini"
	"github.com/codegangsta/martini-contrib/binding"
	_"github.com/go-sql-driver/mysql"
)




func main(){
	m := martini.Classic()
	defer dbmap.Db.Close()

	m.Get("/json_test", func() string {
		return "hello"
	})

	nw := newUser("admADin", "admDDDin", "손건D")
	err := dbmap.Insert(&nw)
	nb := newBus("admin", "hello", "helloworld")
	err = dbmap.Insert(&nb)
	check_err(err, "Insert failed")

	m.Post("/user/login", func() string {
		return "user"
	})

	m.Post("/user/login", binding.Form(user_bind{}), login)


	m.RunOnAddr(":8787")
	m.Run()
}

