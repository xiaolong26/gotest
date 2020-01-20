package test

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"

	//"google.golang.org/appengine/log"
	//"time"
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

	/*var err error
	for i := 1; i <= 64; i <<= 1 {
		DB, err = sql.Open("mysql","gotest:123456@tcp(127.0.0.1:3306)/gotest?charset=utf8" )
		if err != nil {
			//log.Errorf("Connect %s failed: %s; on...: %v","gotest:123456@tcp(127.0.0.1:3306)/gotest?charset=utf8", err, i)
			time.Sleep(time.Second * time.Duration(i))
			continue
		} else {
			break
		}
	}
	DB.SetMaxOpenConns(0)
	DB.SetMaxIdleConns(0)
	for i := 1; i <= 64; i <<= 1 {
		if err := DB.Ping(); err != nil {
			//log.Errorf("Connect %s fatal %s", "gotest:123456@tcp(127.0.0.1:3306)/gotest?charset=utf8", err)
			time.Sleep(time.Second * time.Duration(i))
			continue
		} else {
			break
		}
	}*/
	}

func Query(DB *sql.DB) (int,string,string) {
	var user User
	rows,err := DB.Query("select id,username,passwd from user ")
	if err != nil{
		fmt.Println(err)
	}
	for rows.Next(){
		rows.Scan(&user.Id,&user.Username,&user.Passwd)
		//fmt.Println(id,name,age)
	}
	rows.Close()
	return user.Id,user.Username,user.Passwd
}
