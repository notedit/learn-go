package main

import (
    "fmt"
    "log"
    "strings"
    "database/sql"
    _ "github.com/lxn/go-pgsql"
)


func main(){
    db,err := sql.Open("postgres","sslmode=disable user=user password=password dbname=database")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    
    r,err := db.Query("SELECT a FROM test3 WHERE id = 1;")
    if err != nil {
        log.Print(err)
    }
    
    if !r.Next() {
        log.Print("not next")
    }

    
    _,err = db.Exec("INSERT INTO test4 (a) VALUES (4)")

    if err != nil {
        if strings.Contains(err.Error(),"test4_a_key") {
            fmt.Println("test4_a_key")
        }
    }
}
