package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// User 用户表
type User struct {
	ID   int64
	Name sql.NullString `gorm:"default:'小王子'"`
	Age  int64
}

func main() {
	// mysql数据库连接
	db, err := gorm.Open("mysql", "root:root1234@(127.0.0.1:13306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
	// 同步表字段配置
	db.AutoMigrate(&User{})
	// 创建
	u := User{Name: sql.NullString{String: "", Valid: true}, Age: 18} // z在代码创建User

	db.NewRecord(&u) // 判断主键是否为空

	db.Debug().Create(&u)
	fmt.Println(db.NewRecord(&u))

	// u1 := User{Name: "qimi", Age: 18, Active: true}
	// db.Create(&u1)
	// u2 := User{Name: "jinzhu", Age: 20, Active: false}
	// db.Create(&u2)

	// 5.查询
	var user User
	db.First(&user)

	// 6.更新
	// user.Name = "七米"
	// user.Age = 99
	// db.Debug().Save(&user)
	db.Debug().Model(&user).Update("name", "Simonbin")

	m1 := map[string]interface{}{
		"name":   "Simon-bin",
		"age":    28,
		"active": true,
	}

	db.Debug().Model(&user).Update(m1)                // 更新m1列出来的所有字段
	db.Debug().Model(&user).Select("age").Update(m1)  // 只更新age字段
	db.Debug().Model(&user).Omit("active").Update(m1) // 排除active字段

	rowNum := db.Model(User{}).Update(User{
		Age: 18,
	}).RowsAffected
	fmt.Println(rowNum) // 影响的行数

	db.Delete(&u)
	db.Unscoped().Delete(&u)
}
