package main

import (
	"database/sql"
	"log"
	"time"

	"github.com/coopernurse/gorp"
	_ "github.com/go-sql-driver/mysql"
)

type USER_DB struct {
	Id       int64 `db:"Idx"`
	Created  int64
	UserId   string
	UserPw   string
	UserName string
	SUBJECT  int
	GRADE    int
	CLASS    int
	NUM      int
	authenticated bool `form:"-" db:"-"`
}
func (u *USER_DB) Login() {
	// Update last login time
	// Add to logged-in user's list
	// etc ...
	u.authenticated = true
}

// Logout will preform any actions that are required to completely
// logout a user.
func (u *USER_DB) Logout() {
	// Remove from logged-in user's list
	// etc ...
	u.authenticated = false
}

func (u *USER_DB) IsAuthenticated() bool {
	return u.authenticated
}

func (u *USER_DB) UniqueId() interface{} {
	return u.Id
}

// GetById will populate a user object from a database model with
// a matching id.
func (u *USER_DB) GetById(id interface{}) error {
	err := dbmap.SelectOne(u, "SELECT * FROM USER WHERE UserId = $1", id)
	if err != nil {
		return err
	}

	return nil
}




type BUS struct {
	Id      int64
	Created int64
	Writer  string
	Title   string
	Content string
}

func make_dbmap() *gorp.DbMap {
	db, err := sql.Open("mysql", "tester:tester@tcp(127.0.0.1:3306)/TEST")
	check_err(err, "db connection error")
	log.Println("db connection Ok")

	dialect := gorp.MySQLDialect{"InnoDB", "UTF8"}
	dbmap := &gorp.DbMap{Db: db, Dialect: dialect}

	AddTable(dbmap, USER_DB{}, "USER")
	AddTable(dbmap, BUS{}, "BUSBOARD")
	log.Println("Add Table in gorp Ok")
	return dbmap
}

func AddTable(dbmap *gorp.DbMap, it interface{}, name string) {
	dbmap.AddTableWithName(it, name).SetKeys(true, "Id")
	err := dbmap.CreateTablesIfNotExists()
	check_err(err, "Create tables failed")
}

func newUser(id, pw, name string) USER_DB {
	return USER_DB{
		Created:  time.Now().Unix(),
		UserId:   id,
		UserPw:   hasher(pw),
		UserName: name,
	}
}

func newBus(write, title, content string) BUS {
	return BUS{
		Created: time.Now().Unix(),
		Writer:  write,
		Title:   title,
		Content: content,
	}
}

func selectUser(userID string) USER_DB {
	var user USER_DB
	err := dbmap.SelectOne(&user, "select * from USER where UserId=?", userID)
	check_err(err, "User Select Error")
	return user
}
