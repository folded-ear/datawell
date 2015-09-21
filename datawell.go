package main

import (
    "fmt"
    "github.com/jinzhu/gorm"
)

func main() {
    fmt.Printf("hello, world!  here's a gorm error: %v\n", gorm.InvalidSql)
}
