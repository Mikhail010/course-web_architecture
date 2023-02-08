package mongo

import "github.com/mikhail010/course-web_architecture"

type Db map[int]architecture.Person

func (db Db) Save(idx int, p architecture.Person) {
	db[idx] = p
}

func (db Db) Retrieve(idx int) architecture.Person {
	return db[idx]
}
