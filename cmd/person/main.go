package main

import (
	"fmt"
	"github.com/Mikhail010/course-web_architecture/db_interface"
	"github.com/Mikhail010/course-web_architecture/db_interface/storage/mongo"
	"github.com/Mikhail010/course-web_architecture/db_interface/storage/postgres"
)

func main() {

	mdb := mongo.Db{}
	pgdb := postgres.Db{}

	p1 := db_interface.Person{name: "mika"}
	p2 := db_interface.Person{name: "hector"}

	ps := db_interface.NewPersonService()

	db_interface.Put(mdb, 1, p1)
	db_interface.Put(mdb, 2, p2)
	fmt.Println(db_interface.Get(mdb, 1))
	fmt.Println(get(mdb, 2))

	fmt.Println(ps.GetPerson(1))
	fmt.Println(ps.GetPerson(3))

	db_interface.Put(pgdb, 1, p1)
	db_interface.Put(pgdb, 2, p2)
	fmt.Println(db_interface.Get(pgdb, 1))
	fmt.Println(db_interface.Get(pgdb, 2))
}
