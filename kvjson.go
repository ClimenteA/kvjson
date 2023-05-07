package kvjson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type DB struct {
	path string
}

func New(dbPath string) DB {
	err := os.MkdirAll(dbPath, os.ModePerm)
	if err != nil {
		panic(err)
	}
	return DB{path: dbPath}
}

func (db DB) Set(key string, val interface{}) error {
	entry := filepath.Join(db.path, key+".json")
	file, err := os.Create(entry)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(val)
	if err != nil {
		return err
	}

	return nil
}

func (db DB) Get(key string, val interface{}) error {
	entry := filepath.Join(db.path, key+".json")
	file, err := os.Open(entry)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&val)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (db DB) Del(key string) error {
	entry := filepath.Join(db.path, key+".json")
	err := os.Remove(entry)
	if err != nil {
		return err
	}
	return nil
}
