package main

import (
	"log"

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

	//nw := newUser("jen6", "abcd", "손건")
	//	err := dbmap.Insert(&nw)
	nb := newBus("jen6", "hello", "helloworld")
	err := dbmap.Insert(&nb)
	check_err(err, "Insert failed")

	//	m.Post("/user/login", func() string {
	//		return "use form!"
	//	})
	m.Post("/user/login", binding.Bind(user_bind{}), func(lg user_bind) string {
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
	})
	m.RunOnAddr(":8989")
}
