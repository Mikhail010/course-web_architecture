package main

import (
	"errors"
	"fmt"
)

// 1. Declare entity type
// 2. Declare DB types (mongo and pg)
// 3. Declare storage interface w/methods save and retrieve
// 4. Implement storage interface for mongoDB and postgresDB types
// 5. Create high level funcs (put and get) that use storage interface methods

type person struct {
	name string
}

type mongoDB map[int]person
type postgresDB map[int]person

type storage interface {
	save(idx int, p person)
	retrieve(idx int) person
}

func (m mongoDB) save(idx int, p person) {
	m[idx] = p
}

func (m mongoDB) retrieve(idx int) person {
	return m[idx]
}

func (pg postgresDB) save(idx int, p person) {
	pg[idx] = p
}

func (pg postgresDB) retrieve(idx int) person {
	return pg[idx]
}

type personService struct {
	s storage
}

func (ps *personService) getPerson(idx int) (person, error) {
	p := ps.s.retrieve(idx)
	if p.name == "" {
		return person{}, errors.New(fmt.Sprintf("No person found at %d", idx))
	}
	return p, nil
}

func put(s storage, idx int, p person) {
	s.save(idx, p)
}

func get(s storage, idx int) person {
	return s.retrieve(idx)
}

func main() {
	mdb := mongoDB{}
	pgdb := postgresDB{}

	p1 := person{name: "mika"}
	p2 := person{name: "hector"}

	ps := personService{s: mdb}

	put(mdb, 1, p1)
	put(mdb, 2, p2)
	fmt.Println(get(mdb, 1))
	fmt.Println(get(mdb, 2))

	fmt.Println(ps.getPerson(1))
	fmt.Println(ps.getPerson(3))

	put(pgdb, 1, p1)
	put(pgdb, 2, p2)
	fmt.Println(get(pgdb, 1))
	fmt.Println(get(pgdb, 2))
}
