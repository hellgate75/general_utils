package config

import (
	"fmt"
	"os"
	"testing"
)

func TestReadDatabaseConfig(t *testing.T) {
	var arDbPath string = fmt.Sprintf("%s%c%s%c%s%c%s", "..", os.PathSeparator, "test", os.PathSeparator, "resources", os.PathSeparator, "sample-db.a")
	var dbData map[string][]byte
	var err error

	err, dbData = ReadDatabaseConfig(arDbPath)
	if err != nil {
		t.Fatal("Failed Reading the database : ", err.Error())
	} else if dbData == nil {
		t.Fatal("Database content has not been loaded!!")
	}
}
