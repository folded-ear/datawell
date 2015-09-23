package main

import (
    "fmt"
    "github.com/jinzhu/gorm"
    "database/sql"
    "time"
)

func main() {
    config, err := newConfig()
    if err != nil {
        fmt.Printf("error reading config: %v\n", err)
        return
    }

    fmt.Printf("hello, world!  here's a gorm error: %v.  And here's a config flag: %v\n", gorm.InvalidSql, config.env)

    fmt.Printf("driver: %v, open: %v\n", config.driverName, config.dataSourceName)

    db, err := sql.Open(config.driverName, config.dataSourceName)
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
