package main

import (
	"database/sql"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// User 定义模型
type User struct {
	gorm.Model
	Name         string
	Age          sql.NullInt64
	Birthday     time.Time
	Email        string `gorm:"type:varchar(100);unique_index"`
	Role         string `gorm:"size:255"`        // 设置字段大小为255
	MemberNumber string `gorm:"unique;not null"` // 设置会员(member number) 唯一并且不为空
	Num          int    `gorm:"AUTO_INSCREMENT"` // 设置num 为自增类型
	Address      string `gorm:"index:addr"`      // 给address字段创建名为addr的索引
	IgnoreMe     int    `gorm:"-"`               // 忽略本字段
}

// TableName 将 User 的表名设置为 `profiles`
// func (User) TableName() string {
// 	return "profiles"
// }
// TableName 将 User 的表名设置为 `profiles`
func (u User) TableName() string {
	if u.Role == "admin" {
		return "admin_users"
	}
	return "users"
}

// Animal 动物
type Animal struct {
	AnimalID int64 `gorm:"primary_key"` // 使用AnimalID 做主键
	Name     string
	Age      int64
}

func main() {
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "prefix_" + defaultTableName
	}

	// 连接mysql 数据库
	db, err := gorm.Open("mysql", "root:root1234@(127.0.0.1:13306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	db.AutoMigrate(&User{}, &Animal{})
	// 禁用默认表名的复数形式
	// db.SingularTable(true)

	// 使用User结构体创建 叫 xiaowangzi 的表
	// db.Table("xiaowangzi").CreateTable(&User{})

}
