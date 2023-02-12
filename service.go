package architecture

import (
	"errors"
	"fmt"
)

// Person schema definition to store in a storage
type Person struct {
	Name string
}

// storage actions to read/write persons.
// When retrieving a person, if they don't exist, returns the zero value
type storage interface {
	Save(idx int, p Person)
	Retrieve(idx int) Person
}

// getPersonService gives access to GetPerson action with only what it needs to know (the storage)
type getPersonService struct {
	s storage
}

func NewPersonService(s storage) *getPersonService {
	return &getPersonService{s: s}

}

func (ps *getPersonService) GetPerson(idx int) (Person, error) {
	p := ps.s.Retrieve(idx)
	if p.Name == "" {
		return Person{}, errors.New(fmt.Sprintf("No Person found at %d", idx))
	}
	return p, nil
}

func Put(s storage, idx int, p Person) {
	s.Save(idx, p)
}

func Get(s storage, idx int) Person {
	return s.Retrieve(idx)
}
