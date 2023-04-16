package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// All key and values are being stored as a map
var kvStore = make(map[string]string)

// Coloured logger for error
var Error = log.New(os.Stdout, "\u001b[31mERROR: \u001b[0m", log.LstdFlags|log.Lshortfile)

// Error handling function
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Function to read file
func readFile() {
	file, err := os.ReadFile("store.json")
	check(err)
	err = json.Unmarshal(file, &kvStore)
	check(err)
}

// Function to get value of a key give by user
func get(key string) (string, error) {
	readFile()
	value, ok := kvStore[key]
	if ok == false {
		Error.Println("Key not found")
		return "Key not found", errors.New("not found")
	}
	return value, nil
}

// function to update the key value input from user
// The values are stored in a file
func put(key, value string) (string, error) {
	readFile()
	kvStore[key] = value
	marshalledFile, err := json.Marshal(kvStore)
	check(err)
	err = os.WriteFile("store.json", marshalledFile, 644)
	if err != nil {
		return "Update Failed", errors.New("update failed")
	}
	return "updated", nil
}

func main() {
	// check if file exist. If not create the file
	if _, err := os.Stat("store.json"); os.IsNotExist(err) {
		f, err := os.Create("store.json")
		check(err)
		_, err = f.WriteString("{}")
		check(err)
		defer f.Close()
	}

	// Router for gin web framework
	router := gin.Default()

	// Initialising GET method for key
	router.GET("/key", func(c *gin.Context) {
		key := c.Query("key")
		value, error := get(key)
		if error != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Key not found",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"value": value,
		})
	})

	// Initialising PUT method for key and value
	router.PUT("/key", func(c *gin.Context) {
		key := c.Query("key")
		value := c.Query("value")
		log.Println(key, value)
		value, error := put(key, value)
		if error != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Update failed",
			})
			return
		}
		c.JSON(http.StatusAccepted, gin.H{
			"value": value,
		})
	})

	router.Run()

}
