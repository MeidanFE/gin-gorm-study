package main

import (
	"fmt"
	"os"

	_ "github.com/jinzhu/gorm/dialects/mysql"

	"lession19/dao"
	"lession19/model"
	"lession19/routers"
	"lession19/setting"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage：./lession19 conf/config.ini")
		return
	}

	// 加载配置文件
	if err := setting.Init(os.Args[1]); err != nil {
		fmt.Printf("load config from file failed, err:%v\n", err)
		return
	}
	// 创建数据库
	// sql: CREATE DATABASE bubble
	// mysql数据库连接
	err := dao.InitMySQL(setting.Conf.MySQLConfig)
	if err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer dao.Close()
	// 同步表字段配置
	dao.DB.AutoMigrate(&model.Todo{})
	// 连接数据库

	r := routers.SetupRouter()

	if err := r.Run(fmt.Sprintf(":%d", setting.Conf.Port)); err != nil {
		fmt.Printf("server startup failed, err:%v\n", err)
	}
}
