package architecture

import (
	"fmt"
	"testing"
)

type mockDb map[int]Person

func (db mockDb) Save(idx int, p Person) {
	db[idx] = p
}

func (db mockDb) Retrieve(idx int) Person {
	return db[idx]
}

func TestPut(t *testing.T) {
	mDb := mockDb{}
	expected := Person{Name: "Mika"}

	Put(mDb, 1, expected)

	got := mDb.Retrieve(1)

	if got != expected {
		t.Fatalf("Expected %v, got %v", expected, got)
	}
}

func ExamplePut() {
	mDb := mockDb{}
	expected := Person{Name: "Mika"}

	Put(mDb, 1, expected)
	got := mDb.Retrieve(1)
	fmt.Println(got)
	// Output: {Mika}
}
