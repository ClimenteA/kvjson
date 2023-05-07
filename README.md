# KVJson
Simple key/value store for Go using plain json files.


## Install

```bash
go get -u github.com/ClimenteA/kvjson
```

## Usage

```go
package main

import "github.com/ClimenteA/kvjson"


func main() {

	type Person struct {
		Name string
		Age  int
	}

	type Car struct {
		Model string
		Year  int
	}

    // Next, in the main func initialize the "db" with a folder path
	// Use `kvjson.DB` to get the types. Ex: `func someFunc(db kvjson.DB) {etc}` 
	db := kvjson.New("./db")

    // Set a `key` with an initialized `struct`
    // In the `./db` folder a `key.json` will be saved
	db.Set("person", Person{Name: "Alin", Age: 30})
	db.Set("car", Car{Model: "Dacia", Year: 2020})

    // Get value of a `key`
    // This will just unmarshal the data into the struct
	var car Car
	err := db.Get("car", &car)
	if err != nil {
		panic(err)
	}
	fmt.Println(car)

	var person Person
	err = db.Get("person", &person)
	if err != nil {
		panic(err)
	}
	fmt.Println(person)

    // Delete a `key`
	err = db.Del("person")
    if err != nil {
		panic(err)
	}
	err = db.Del("car")
    if err != nil {
		panic(err)
}
```


