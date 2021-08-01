package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	// 包的 init 函数会注册 sqlserver 驱动
	_ "github.com/denisenkom/go-mssqldb"
)

// sql.DB 用来操作数据库，它代表了 0 或多个底层连接的池，这些连接由 sql 包来维护，sql 包会字段创建和释放这些连接
// - 它是线程安全的
// - sql.DB 不需要进行关闭 (当然如果想自己关闭也可以)
// - 它是用来处理数据库的，而不是实际的连接
// - 这个抽象包含了数据库连接的池，而且会对此进行维护
// - 使用 sql.DB 时，可以定义他的全局变量进行使用，也可以将它传递到函数里
// sql.DB 类型上用于查询的方法有
// - Query/QueryContext: 返回多条
// - QueryRow/QueryRowContext: 返回一条
// sql.DB 类型上用于更新的方法有
// - Exec/ExecContext
var db *sql.DB

const (
	server   = "xxx.database.windows.net"
	port     = 1433
	user     = "xxxx"
	password = "xxxx"
	database = "go-db"
)

func main() {
	var err error

	// 连接数据库
	connStr := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", server, user, password, port, database)
	/**
	 * Open 函数并不会连接数据库，也不会验证参数，只是把后续连接到数据库所必须的 struct 给设置好了, 真正的连接是在被需要的时候才进行懒设置的
	 */
	db, err = sql.Open("sqlserver", connStr)
	if err != nil {
		log.Fatalln(err.Error())
	}

	/**
	 * PingContext 用来验证与数据库的连接是否仍然有效，如有必要的话则建立一个连接
	 */
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println("Connected")

	// 查询一个
	one, err := getOne(103)
	if err == nil {
		log.Fatal(err.Error())
	}
	fmt.Println(one)

	// 查询多个
	apps, err := getMany(103)
	if err == nil {
		log.Fatal(err.Error())
	}
	fmt.Println(apps)

	// 更新
	one.name += " 1234"
	one.order++
	err = one.Update()
	if err == nil {
		log.Fatal(err.Error())
	}
	one2, _ := getOne(103)
	fmt.Println(one2)

	// 删除
	// one2.Delete()

	// 创建
	a := app{
		name:   "Test",
		order:  123,
		level:  1,
		status: 1,
	}
	err = a.Insert()
	if err == nil {
		log.Fatal(err.Error())
	}
	one3, _ := getOne(a.ID)
	fmt.Println(one3)
}
