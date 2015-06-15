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

	m.Post("/user/login", binding.Bind(user_bind{}), func(lg user_bind) string {
		return login(lg)
	})

	m.RunOnAddr(":8989")
}
