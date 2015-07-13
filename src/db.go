package main

import (
	"database/sql"
	"log"
	"time"
	"gopkg.in/gorp.v1"
	_ "github.com/go-sql-driver/mysql"
	"errors"
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

func (u* USER_DB) insert() {
	err := dbmap.Insert(u);
	check_err(err, "error in insert user")
}

type Board interface {
	search(bf Board_find) []interface{}
	list(idx int) ([]interface{}, error)
	write(T ANY)
	update(T ANY)
}

type BUS struct {
	Id       int64    `db:"Idx"`
	Created  int64
	Writer   string    `db:"Writer"`
	WriterId int64
	Title    string    `db:"Title"`
	Content  string    `db:"Content"`
	Want     int64
	Status   int64
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

func (b *BUS) view(id int64) {
	err := dbmap.SelectOne(b, "select * from BUSBOARD where Idx = ?", id)
	check_err(err, "error in view bus")
}

func (b* BUS) write() {
	err := dbmap.Insert(b)
	check_err(err, "error in bus write")
}

func calc_limitPage(onepage int64, count int64, idx int64) error {
	var total_page int64
	if idx == 0 {
		return errors.New("invaild busboard idx")
	}

	if count%onepage == 0 {
		total_page = count/onepage
	} else {
		total_page = count/onepage + 1
	}

	if idx > total_page || idx < 0 {
		return errors.New("invaild busboard idx")
	}
	return nil
}

func (b BUS) list() ([]BUS, error) {
	//	TODO 나중에 리스트 형식으로 다시 구현하기..
	//	const onePage int64 = 5
	//	idx64 := int64(idx)
	//	count, err := dbmap.SelectInt("select count(*) from BUSBOARD where Status = 0")
	//	check_err(err, "error in count busboard")
	//	err = calc_limitPage(onePage, count, idx64)
	//	if err != nil {
	//		return []BUS{}, errors.New("invaild idx")
	//	}
	var arr []BUS
	_, err := dbmap.Select(&arr, "select * from BUSBOARD where Status = 0 order by Created desc")
	if err != nil {
		return []BUS{}, errors.New("fail in select busboard")
	}
	return arr, nil
}

func (b * BUS) update() {
	_, err := dbmap.Update(b)
	check_err(err, "error in bus update")
}

type PROFILE struct {
	Id    int64
	Best  string
	Can   string
	Intro string
	Image string
}

func (p * PROFILE) Get(WriterIdx int64) error {
	err := dbmap.SelectOne(p, "SELECT * FROM PROFILE WHERE Id = ?", WriterIdx)
	if err != nil {
		return err
	}
	return nil
}

//type Board interface {
//	search(bf Board_find) []interface{}
//	list(idx int) ([]interface{}, error)
//	write(T ANY)
//	update(T ANY)
//}

type PORTFOLIO struct {
	Id       int64
	Created  int64
	WriterId int64
	Writer   string
	Content  string
	Image    string
}

func (pf * PORTFOLIO) write() {
	err := dbmap.Insert(pf)
	check_err(err, "error in insert portfolio")
}

func (pf * PORTFOLIO) view(id int64) {
	err := dbmap.SelectOne(pf, "SELECT * FROM PORTFOLIO WHERE Id = ?", id)
	check_err(err, "error in view porfol")
}



func make_dbmap() *gorp.DbMap {
	db, err := sql.Open("mysql", "tester:tester@tcp(127.0.0.1:3306)/TEST")
	check_err(err, "db connection error")
	log.Println("db connection Ok")

	dialect := gorp.MySQLDialect{"InnoDB", "UTF8"}
	dbmap_buf := &gorp.DbMap{Db: db, Dialect: dialect}

	AddTable(dbmap_buf, true, USER_DB{}, "USER")

	table := AddTable(dbmap_buf, true, BUS{}, "BUSBOARD")
	table.ColMap("Writer").SetMaxSize(10)
	table.ColMap("Title").SetMaxSize(25)
	table.ColMap("Content").SetMaxSize(50)

	AddTable(dbmap_buf, false, PROFILE{}, "PROFILE")
	table = AddTable(dbmap_buf, true, PORTFOLIO{}, "PORTFOLIO")
	table.ColMap("Writer").SetMaxSize(10)
	log.Println("Add Table in gorp Ok")

	return dbmap_buf
}

func AddTable(dbmap_buf *gorp.DbMap, auto_inc bool, it interface{}, name string) *gorp.TableMap {
	var table *gorp.TableMap = dbmap_buf.AddTableWithName(it, name).SetKeys(auto_inc, "Id")
	err := dbmap_buf.CreateTablesIfNotExists()
	check_err(err, "Create tables failed")
	return table
}

func newUser(id, pw, name string) USER_DB {
	return USER_DB{
		Created:  time.Now().UnixNano(),
		UserId:   id,
		UserPw:   hasher(pw),
		UserName: name,
	}
}

func newBus(idx, want int64, write, title, content string) BUS {
	return BUS{
		Id:0,
		Created: time.Now().UnixNano(),
		Writer:  write,
		WriterId:idx,
		Title:   title,
		Content: content,
		Want: want,
	}
}

func newPorfolio(WriterId int64, writer, content, image string) PORTFOLIO {
	return PORTFOLIO{
		Id:0,
		Created: time.Now().UnixNano(),
		Writer: writer,
		WriterId: WriterId,
		Content: content,
		Image: image,
	}
}

func selectUser(userID string) USER_DB {
	var user USER_DB
	err := dbmap.SelectOne(&user, "select * from USER where UserId=?", userID)
	check_err(err, "User Select Error")
	return user
}