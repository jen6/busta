package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"log"
	"fmt"
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
	//	ud := newUser("test1", "test1", "박선현")
	//	ud.insert()
	//	ud = newUser("test2", "test2", "권욱제")
	//	ud.insert()
	//	ud = newUser("test3", "test3", "강명서")
	//	ud.insert()
}

func substring(s string, len int) string {
	by := []byte(s)

	if int(by[len - 1]) >= 224 {
		len += 2
	} else if int(by[len - 1]) >= 192 && int(by[len - 1]) < 224 {
		len += 1
	}

	return s[0:len]
}


type ANY interface{}