package storage

import "github.com/mikhail010/go/db_interface"

type DB map[int]db_interface.Person

func (db DB) save(idx int, p db_interface.Person) {
	db[idx] = p
}

func (db DB) retrieve(idx int) db_interface.Person {
	return db[idx]
}
