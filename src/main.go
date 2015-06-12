package main

import (
	"github.com/codegangsta/martini-contrib/binding"
	"github.com/go-martini/martini"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
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

	//	m.Post("/user/login", func() string {
	//		return "use form!"
	//	})
	m.Post("/user/login", binding.Form(user_bind{}), login)

	m.RunOnAddr(":8989")
	m.Run()
}
