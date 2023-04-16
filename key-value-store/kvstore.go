package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var kvStore = make(map[string]string)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func get(key string) string {
	value, ok := kvStore[key]
	if ok == false {
		return "not found"
	}
	return value
}
func put(key, value string) {
	kvStore[key] = value
}

func main() {
	// check if file exist. If not create it
	if _, err := os.Stat("kv.json"); os.IsNotExist(err) {
		f, err := os.Create("kv.json")
		check(err)
		_, err = f.WriteString("{}")
		check(err)
		defer f.Close()
	}

	file, err := os.ReadFile("kv.json")
	check(err)

	json.Unmarshal(file, &kvStore)

	value := get("hello")
	fmt.Println(value)

	put("tax", "submitted")

	value2 := get("tax")
	fmt.Println(value2)
}
