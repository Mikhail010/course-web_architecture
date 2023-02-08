package architecture

import (
	"errors"
	"fmt"
)

// 1. Declare entity type
// 2. Declare DB types (mongo and pg)
// 3. Declare storage interface w/methods save and retrieve
// 4. Implement storage interface for mongoDB and postgresDB types
// 5. Create high level funcs (put and get) that use storage interface methods

type Person struct {
	name string
}

type storage interface {
	save(idx int, p Person)
	retrieve(idx int) Person
}

type personService struct {
	s storage
}

func NewPersonService(s storage) *personService {
	return &personService{s: s}

}

func (ps *personService) GetPerson(idx int) (Person, error) {
	p := ps.s.retrieve(idx)
	if p.name == "" {
		return Person{}, errors.New(fmt.Sprintf("No Person found at %d", idx))
	}
	return p, nil
}

func Put(s storage, idx int, p Person) {
	s.save(idx, p)
}

func Get(s storage, idx int) Person {
	return s.retrieve(idx)
}
