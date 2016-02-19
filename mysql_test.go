package mysql_cache_performance

import (
 "testing"
 "database/sql"
 "math/rand"

 _ "github.com/go-sql-driver/mysql"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}

func BenchmarkMySqlGet(b *testing.B) {
    var id int64
    db, err := sql.Open("mysql", "user:pass@tcp(mysql-master:3306)/cached_db?autocommit=true")
    if err != nil {
        b.Fatalf("couldn't connect to mysql")
    }
    stmtInsert, err := db.Prepare("INSERT INTO cached (data) VALUES( ? )")
    if err != nil {
        panic(err.Error())
    }
    data, err := stmtInsert.Exec(randSeq(30))
    if err != nil {
        panic(err.Error())
    }
    id, err = data.LastInsertId()
    if err != nil {
        panic(err.Error())
    }
    dbr, err := sql.Open("mysql", "user:pass@tcp(mysql-slave:3306)/cached_db?autocommit=true")
    if err != nil {
        b.Fatalf("couldn't connect to mysql")
    }

    stmtGet, err := dbr.Prepare("SELECT data FROM cached WHERE id = ?")
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }

    //make sure it's on reader
    var scanData string
    for {
        err = stmtGet.QueryRow(id).Scan(&scanData)
        if err == nil {
            break
        }
    }

    for i := 0; i < b.N; i++ {
        err = stmtGet.QueryRow(id).Scan(&scanData)
        if err != nil {
            panic(err.Error()) // proper error handling instead of panic in your app
        }
    }
}

