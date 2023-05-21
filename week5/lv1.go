package main

import (
	"csa/model"
	"database/sql" //标准库
	"fmt"
	_ "github.com/go-sql-driver/mysql" //我们使用的mysql，需要导入相应驱动包，否则会报错
	"log"
)

// 定义一个全局对象db
var db *sql.DB

func initDB() {
	var err error
	// 设置一下dns charset:编码方式 parseTime:是否解析time类型 loc:时区
	dsn := "root:277187@tcp(127.0.0.1:3306)/student?charset=utf8mb4&parseTime=True&loc=Local"
	// 打开mysql驱动
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("DB connect success")
	return
}

// 查询多条数据示例
func queryMultiRowDemo() {
	sqlStr := "select id, name,sex, age from student where id > ?"
	rows, err := db.Query(sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	// 非常重要：关闭rows释放持有的数据库链接
	defer rows.Close()

	// 循环读取结果集中的数据
	for rows.Next() {
		var u model.Student
		err := rows.Scan(&u.Id, &u.Name, &u.Sex, &u.Age)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("id:%d name:%s sex:%s age:%s\n", u.Id, u.Name, u.Sex, u.Age)
	}
}

// InsertStudent 插入数据
func InsertStudent(st model.Student) {
	sqlStr := "insert into student(name,sex,age) values (?,?,?)"
	_, err := db.Exec(sqlStr, st.Name, st.Sex, st.Age)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	log.Println("insert success")
}

func main() {
	//初始化连接
	initDB()
	var stu [10]model.Student
	stu = [10]model.Student{
		{1, "小红", "女", "19"},
		{2, "小王", "男", "16"},
		{3, "小黑", "男", "17"},
		{4, "小x", "男", "17"},
		{5, "小y", "女", "17"},
		{6, "小a", "男", "21"},
		{7, "小b", "女", "16"},
		{8, "小c", "女", "13"},
		{9, "小d", "女", "19"},
		{10, "小e", "男", "17"},
	}
	for i := 0; i < 10; i++ {
		InsertStudent(stu[i])
	}
	queryMultiRowDemo()
}
