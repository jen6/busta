package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"log"

)

func init() {
	dbmap = make_dbmap()
}

func hasher(str string) string {
	it := sha256.New()
	hash_arr := []byte(str)
	copy(hash_arr[:], str)
	it.Write(hash_arr)
	md := it.Sum(nil)
	mdStr := hex.EncodeToString(md)
	return mdStr
}

func check_err(err error, msg string) {
	if err != nil {
		log.Println(msg, err)
	}
}


func struct2json(it interface{}) string {
	b, _ := json.Marshal(it)
	return string(b)
}

func addmysqldata() {
	ui := user_info{
		UserName:"손건",
	}
	ud := USER_DB{}
	ud.search_one(ui)
	if ud.Id == 0 {
		ud = newUser("jen6", "abcd", "손건")
		ud.insert()
	}

}

type ANY interface{}