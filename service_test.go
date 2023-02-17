package architecture

import (
	"fmt"
	"github.com/golang/mock/gomock"
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
	ctl := gomock.NewController(t)
	expected := Person{Name: "Mika"}

	mDb := NewMockstorage(ctl)
	mDb.EXPECT().Save(1, expected).MinTimes(1).MaxTimes(1)

	Put(mDb, 1, expected)

	ctl.Finish()
}

func ExamplePut() {
	mDb := mockDb{}
	expected := Person{Name: "Mika"}

	Put(mDb, 1, expected)
	got := mDb.Retrieve(1)
	fmt.Println(got)
	// Output: {Mika}
}
