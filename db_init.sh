#!/bin/zsh

# 初始化环境
docker run --name mysql8091 -p 13306:3306 -e MYSQL_ROOT_PASSWORD=root1234 -d mysql:8.0.19

#新建终端建立连接
docker run -it --network host --rm mysql mysql -h127.0.0.1 -P13306 --default-character-set=utf8mb4 -uroot -proot1234

# 创建数据库
create database db1;

# 查看所有数据库
show databases;

# 使用数据库db1
use db1;

# 查看所有的表
show tables;