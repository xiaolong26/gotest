package db

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id int `db:"id"`
	Username string  `db:"username"`  //由于在mysql的users表中name没有设置为NOT NULL,所以name可能为null,在查询过程中会返回nil，如果是string类型则无法接收nil,但sql.NullString则可以接收nil值
	Passwd string `db:"string"`
}

var DB *sql.DB

func init(){
	var err error
	DB, err = sql.Open("mysql", "gotest:123456@tcp(127.0.0.1:3306)/gotest?charset=utf8")
	if err != nil{
		fmt.Println(err)
	}
}

func Checker(c *User) map[string]string {
	var user User
	rows,err := DB.Query("select id,username,passwd from user ")
	if err != nil{
		fmt.Println(err)
	}
	msg := make(map[string]string)
	for rows.Next(){
		rows.Scan(&user.Id,&user.Username,&user.Passwd)
		if user.Username == c.Username && user.Passwd == c.Passwd{
			msg["Msg"] =  "login"
			break
		}else{
			msg["Msg"] = "not login"
			errors.New()
		}
	}
	return msg
}
