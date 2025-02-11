package jsondb_test

import (
	"os"
	"testing"

	"github.com/pnkj-kmr/looking-glass/repos/jsondb"
)

func TestNew(t *testing.T) {
	path := "ipinfo"
	_, err := jsondb.New(path)
	if err == nil {
		t.Error("Test failed - ", err)
	}
}

func TestGetAll(t *testing.T) {
	// testing empty dir
	db, _ := jsondb.New("ipinfo")

	data := db.GetAll()
	if len(data) != 0 {
		t.Error("Test failed - ", data)
	}
}

func TestGet(t *testing.T) {
	// testing empty dir
	db, _ := jsondb.New("ipinfo")

	_, err := db.Get("ip-dummy")
	if os.IsExist(err) {
		t.Error("Test failed - ", err)
	}
}

func TestInsert(t *testing.T) {
	// testing empty dir
	db, _ := jsondb.New("ipinfo")

	var data []byte
	data = append(data, 99)
	err := db.Insert("ip-dummy", data)
	if err != nil {
		t.Error("Test failed - ", err)
	}
}

func TestGet2(t *testing.T) {
	// testing empty dir
	db, _ := jsondb.New("ipinfo")

	_, err := db.Get("ip-dummy")
	if os.IsNotExist(err) {
		t.Error("Test failed - ", err)
	}
}

func TestDelete(t *testing.T) {
	// testing empty dir
	db, _ := jsondb.New("ipinfo")

	err := db.Delete("test_dummp")
	if os.IsExist(err) {
		t.Error("Test failed - ", err)
	}
}
