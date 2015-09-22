package main

import (
    "fmt"
    "flag"
    "github.com/jinzhu/gorm"
    "github.com/vharitonsky/iniflags"
    "database/sql"
    "time"
    "bitbucket.org/liamstask/goose/lib/goose"
)

var (
    gooseEnv = flag.String("env", "development", "the Goose environment to use")
)

func main() {
    iniflags.Parse()
    fmt.Printf("hello, world!  here's a gorm error: %v.  And here's the flag: %v\n", gorm.InvalidSql, *gooseEnv)

    gDbConf, err := goose.NewDBConf("./db", *gooseEnv, "")
    if err != nil {
        fmt.Printf("error reading config: %v\n", err)
        return
    }
    fmt.Printf("driver: %v, open: %v\n", gDbConf.Driver.Name, gDbConf.Driver.OpenStr)

    db, err := sql.Open(gDbConf.Driver.Name, gDbConf.Driver.OpenStr)
    if err != nil {
        fmt.Printf("connect error: %v\n", err)
        return
    }
    var now time.Time
    err = db.QueryRow("select now()").Scan(&now)
    if err != nil {
        fmt.Printf("query error: %v\n", err)
        return
    }
    fmt.Printf("it's %v\n", now)

}
