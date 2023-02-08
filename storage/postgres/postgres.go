package postgres

import "github.com/mikhail010/go/db_interface"

type Db map[int]db_interface.Person

func (db Db) Save(idx int, p db_interface.Person) {
	db[idx] = p
}

func (db Db) Retrieve(idx int) db_interface.Person {
	return db[idx]
}
