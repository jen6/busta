package main

import (
	"database/sql"
	"log"
	"time"
	"gopkg.in/gorp.v1"
	_ "github.com/go-sql-driver/mysql"
)


type USER_DB struct {
	Id            int64 `db:"Idx"`
	Created       int64
	UserId        string
	UserPw        string
	UserName      string
	SUBJECT       int
	GRADE         int
	CLASS         int
	NUM           int
	authenticated bool `form:"-" db:"-"`
}
//TODO: 여기 주석 정리 해야할듯
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
	log.Println(id);
	err := dbmap.SelectOne(u, "SELECT * FROM USER WHERE Idx = ?", id)
	if err != nil {
		return err
	}

	return nil
}

func (u* USER_DB) search_one(ui User_Interface) {
	query, query_map := ui.Prepare()
	err := dbmap.SelectOne(u, query, query_map)
	if err != nil {
		log.Print(err)
	}
}

func (u* USER_DB) search_arr(ui User_Interface) []USER_DB {
	var arr []USER_DB
	query, query_map := ui.Prepare()
	_, err := dbmap.Select(&arr, query, query_map)
	if err != nil {
		log.Print(err)
	}
	return arr
}

type Board interface {
	search(bf Board_find) []interface{}
	list(idx int) []interface{}
	write(T ANY)
	update(T ANY)
}

type BUS struct {
	Id       int64    `db:"Idx, primarykey, autoincrement"`
	Created  int64
	Writer   string    `db:"Writer,size:10"`
	WriterId int64
	Title    string    `db:"Title,size:50"`
	Content  string    `db:"Content,size:50"`
	Want     int64
	status   int64
}

func (b BUS) search(bf Board_find) []BUS {
	var arr []BUS
	query, query_map := bf.Prepare()
	_, err := dbmap.Select(&arr, query, query_map)
	if err != nil {
		log.Print(err)
	}
	return arr
}

func (b* BUS) write() {
	err := dbmap.Insert(b)
	check_err(err, "error in bus write")
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
	var table *gorp.TableMap = dbmap.AddTableWithName(it, name).SetKeys(true, "Id")
	err := dbmap.CreateTablesIfNotExists()
	check_err(err, "Create tables failed")
	log.Print(table.TableName)
}

func newUser(id, pw, name string) USER_DB {
	return USER_DB{
		Created:  time.Now().UnixNano(),
		UserId:   id,
		UserPw:   hasher(pw),
		UserName: name,
	}
}

func newBus(idx int64, write, title, content string) BUS {
	return BUS{
		Id:0,
		Created: time.Now().UnixNano(),
		Writer:  write,
		WriterId:idx,
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


