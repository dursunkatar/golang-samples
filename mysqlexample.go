package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Animal struct {
	Age  int64
	Name string
}

func main() {
	db, err := gorm.Open("mysql", "root:12345@/db_name?charset=utf8&parseTime=True&loc=Local")
	//var animal = Animal{Age: 28, Name: "fesserf"}
	//db.Create(&animal)
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}
	var data = Animal{}
	db.Take(&data)

	fmt.Println(data.Name)

}
