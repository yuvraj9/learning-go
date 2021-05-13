package greetings
import (
	"fmt"
	"math/rand"
	"time"
	"errors"
)

func Hello(name string) (string , error) {
	if name == "" {
		return "" , errors.New("Name empty!")
	}
	message := fmt.Sprintf(RandomString(), name)
	return message, nil
}
func init() {
    rand.Seed(time.Now().UnixNano())
}

func Hellos(names []string) (map[string]string, error) {
	
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)

		if err != nil {
			return nil, err
		}

		messages[name] = message
	}
	return messages, nil
}

func RandomString() string {
	formats := []string{
		"Hi %v, Welcome!",
		"Hello, Welcome to the team %v",
		"Hail %v and also hydra!",
	}
	return formats[rand.Intn(len(formats))]
}