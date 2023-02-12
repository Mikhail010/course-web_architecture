package main

import (
	"fmt"
	architecture "github.com/mikhail010/course-web_architecture"
	"github.com/mikhail010/course-web_architecture/storage/mongo"
	"github.com/mikhail010/course-web_architecture/storage/postgres"
)

func main() {
	mdb := mongo.Db{}
	pgdb := postgres.Db{}

	p1 := architecture.Person{Name: "mika"}
	p2 := architecture.Person{Name: "hector"}

	ps := architecture.NewPersonService(mdb)

	architecture.Put(mdb, 1, p1)
	architecture.Put(mdb, 2, p2)
	fmt.Println(architecture.Get(mdb, 1))
	fmt.Println(architecture.Get(mdb, 2))

	fmt.Println(ps.GetPerson(1))
	fmt.Println(ps.GetPerson(3))

	architecture.Put(pgdb, 1, p1)
	architecture.Put(pgdb, 2, p2)
	fmt.Println(architecture.Get(pgdb, 1))
	fmt.Println(architecture.Get(pgdb, 2))
}
