package main

import (
	"database/sql"
	errors2 "errors"
	"fmt"
	xerr "github.com/pkg/errors"
)

func main() {
	UserController()
}

//接口层
func UserController() {
	if err := UserDao(); err != nil {
		fmt.Printf("%+v", err)
	}

}

//Dao层
func UserDao() error {
	data, err := FindUser()
	if err != nil {
		if errors2.Is(err, sql.ErrNoRows) {
			return xerr.Wrap(err, "not found user")
		} else {
			return err
		}
	}
	fmt.Println(data)
	return nil
}

//查询用户
func FindUser() (interface{}, error) {
	return nil, sql.ErrNoRows
}
