package main

import (
    "fmt"
    "os"
    "github.com/lxn/go-pgsql"
)

func main() {

    conn,err := pgsql.Connect("dbname=database user=user port=5432",pgsql.LogError)
    
    if err != nil {
        os.Exit(1)
    }
    defer conn.Close()

    rs,err := conn.Query("SELECT 1 AS num; SELECT 2 AS num; SELECT 3 AS num;")
    if err != nil {
        os.Exit(1)
    }
    defer rs.Close()

    for {
        hasRow,err := rs.FetchNext()
        if err != nil {
            os.Exit(1)
        }
        if hasRow {
            num,_,_ := rs.Int(0)
            fmt.Println("num:",num)
        } else {
            hasResult,err := rs.NextResult()
            if err != nil {
                os.Exit(1)
            }
            if !hasResult {
                break
            }
            fmt.Println("next result")
        }
    }

}

