package storage

import "github.com/mikhail010/go/db_interface"

type db map[int]db_interface.Person

func (db db) save(idx int, p db_interface.Person) {
	db[idx] = p
}

func (db db) retrieve(idx int) db_interface.Person {
	return db[idx]
}
