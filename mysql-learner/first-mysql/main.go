package main

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/go-sql-driver/mysql"
)

func main() {
    // 连接到数据库
    db, err := sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/school")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // 插入记录
    insertStmt, err := db.Prepare("INSERT INTO student(name, age, class) VALUES(?, ?, ?)")
    if err != nil {
        log.Fatal(err)
    }
    defer insertStmt.Close()

    for i := 1; i <= 10; i++ {
        _, err := insertStmt.Exec(fmt.Sprintf("Student %d", i), 20+i, fmt.Sprintf("Class %d", i))
        if err != nil {
            log.Fatal(err)
        }
    }

    // 读取并打印记录
    rows, err := db.Query("SELECT id, name, age, class FROM student")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    for rows.Next() {
        var id int
        var name string
        var age int
        var class string
        err = rows.Scan(&id, &name, &age, &class)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Printf("ID: %d, Name: %s, Age: %d, Class: %s\n", id, name, age, class)
    }
}
