package main

import (
	"database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql"   //默认执行init方法
)
// 定义一个全局对象db
var db *sql.DB
//定义一个结构体接收查询的数据
type user_auth struct {
	id int
	user string
	passwd string
}
// 定义一个初始化数据库的函数
func initDB() (err error) {
	// DSN:Data Source Name
	dsn := "root:root@tcp(172.31.230.85:3306)/blog?charset=utf8mb4&parseTime=True"
	// 不会校验账号密码是否正确
	// 注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}
//查询操作
func queryRow()  {
	sqlStr:="select id,username,password from blog_auth where id=?"
	var u user_auth
	err := db.QueryRow(sqlStr,1).Scan(&u.id,&u.user,&u.passwd)
	if err!=nil{
		fmt.Printf("查询失败：%s\n",err)
		return
	}
	fmt.Printf("id:%d,user:%s,passwd:%s\n",u.id,u.user,u.passwd)
	return


}

func queryMultiRowDemo()  {
	sqlstr:="select * from blog_auth where id > ?"
	var u user_auth
	rows,err:=db.Query(sqlstr,0)
	if err!=nil{
		fmt.Printf("query faild err:%s",err)
		return
	}
	defer rows.Close()
	for rows.Next(){
		err=rows.Scan(&u.id,&u.user,&u.passwd)
		if err!=nil{
			fmt.Printf("scan faild :%s",err)
			return
		}
		fmt.Printf("id:%d user:%s passwd:%s\n", u.id, u.user, u.passwd)
	}
	
}
func insertRow()  {
	sqlstr:= "insert into blog_auth(username,password) values (?,?)"
	row,err := db.Exec(sqlstr,"insert","123456")
	if err!=nil{
		fmt.Printf("insert faild :%s",err)
		return
	}
	Id , err :=row.LastInsertId()
	if err!= nil{
		fmt.Printf("get lastId faild :%s",err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", Id)


}
func updateRow()  {
	sqlstr:="update blog_auth set password=? where id=?"
	row,err := db.Exec(sqlstr,"123456789",4)
	if err!=nil{
		fmt.Printf("update faild :%s",err)
		return
	}
	number,err := row.RowsAffected()
	if err !=nil{
		fmt.Printf("get changeRow faild :%s",err)
		return
	}
	fmt.Printf("update success, affected rows:%d\n", number)

}
func delRow()  {
	sqlstr:= "delete from blog_auth where id=?"
	row,err := db.Exec(sqlstr,4)
	if err != nil{
		fmt.Printf("del faild :%s",err)
		return
	}
	number,err:= row.RowsAffected()
	if err != nil {
		fmt.Printf("get del number faild:%s",err)
		return
	}
	fmt.Printf("delete success, affected rows:%d\n", number)

}
func prepareQuery()  {
	sqlstr := "select * from blog_auth where id>?"
	per ,err := db.Prepare(sqlstr)
	if err != nil {
		fmt.Printf("prepare faild :%s",err)
		return
	}
	defer db.Close()
	row,err := per.Query(0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	defer row.Close()
	for row.Next(){
		var u user_auth
		err = row.Scan(&u.id,&u.user,&u.passwd)
		if err != nil {
			fmt.Printf("scan faild :%s",err)
			return
		}
		fmt.Printf("id:%d user:%s passwd:%s\n", u.id, u.user, u.passwd)
	}

}
func transaction()  {
	ts,err := db.Begin()
	if err != nil {
		if ts!=nil {
			ts.Rollback()
		}
		fmt.Printf("begin trans faild :%s",err)
		return
	}
	sqlstr1 := "insert into blog_auth(username,password) values (?,?)"
	row,err := ts.Exec(sqlstr1,"trans","123456")
	if err != nil {
		ts.Rollback()
		fmt.Printf("first command faild :%s",err)
		return
	}
	affRow1, err := row.RowsAffected()
	if err != nil {
		ts.Rollback() // 回滚
		fmt.Printf("exec ret1.RowsAffected() failed, err:%v\n", err)
		return
	}
	sqlStr2 := "Update blog_auth set username='testload' where id=?"
	ret2, err := ts.Exec(sqlStr2, 3)
	if err != nil {
		ts.Rollback() // 回滚
		fmt.Printf("exec sql2 failed, err:%v\n", err)
		return
	}
	affRow2, err := ret2.RowsAffected()
	if err != nil {
		ts.Rollback() // 回滚
		fmt.Printf("exec ret2.RowsAffected() failed, err:%v\n", err)
		return
	}
	fmt.Println(affRow1, affRow2)
	if affRow1==1 && affRow2==1{
		fmt.Printf("开始提交事务！")
		ts.Commit()

	}else {
		ts.Rollback()
		fmt.Printf("事务回滚了！")
		return
	}
	fmt.Printf("transaction success!")
	
}

func main() {
	err := initDB() // 调用输出化数据库的函数
	if err != nil {
		fmt.Printf("init db failed,err:%v\n", err)
		return
	}
	fmt.Println("连接数据库成功！")
	//fmt.Printf("单次查询结果：")
	//queryRow()
	//fmt.Printf("\n多次查询结果：")
	//queryMultiRowDemo()
	//fmt.Printf("\n插入数据结构：")
	//insertRow()
	//updateRow()
	//delRow()
	//fmt.Printf("\n预处理多次查询结果：\n")
	//prepareQuery()
	fmt.Printf("运行事务结果：\n")
	transaction()

}