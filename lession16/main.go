package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// UserInfo --> 数据表
type UserInfo struct {
	ID     uint
	Name   string
	Gender string
	Hobby  string
}

func main() {
	// 连接mysql 数据库
	db, err := gorm.Open("mysql", "root:root1234@(127.0.0.1:13306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	// 创建表, 自动迁移 (把结构体和数据表进行对应)
	db.AutoMigrate(&UserInfo{})

	// 创建数据行
	// u1 := UserInfo{3, "Simon-bin", "男", "双色球"}
	// db.Create(&u1)
	// 查询
	var u []UserInfo

	db.Find(&u, &UserInfo{Hobby: "双色球"})
	fmt.Printf("u:%#v\n", u)
	// 更新
	// db.Model(&u).Update("hobby", "双色球")
	// 删除
	// db.Delete(&u)
}
