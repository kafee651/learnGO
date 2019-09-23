package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type animal struct {
	id         int
	animalType string
	nickname   string
	zone       int
	age        int
}

func main() {
	db, err := sql.Open("mysql", "root:jo5kGDw5LJEpRt@tcp(35.193.220.71:3306)/Dino")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("Select * from Dino.animals where id=1")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	animals := []animal{}
	for rows.Next() {
		a := animal{}
		err := rows.Scan(&a.id, &a.animalType, &a.nickname, &a.zone, &a.age)
		if err != nil {
			log.Println(err)
			continue
		}
		animals = append(animals, a)
	}
	fmt.Println(animals)
}
