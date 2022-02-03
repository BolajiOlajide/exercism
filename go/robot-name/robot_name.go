package robotname

import (
	"errors"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// Define the Robot type here.

// Robot representation of the robot we're designing
type Robot struct {
	name string
}

var random = rand.New(rand.NewSource(time.Now().UnixNano()))
var charset = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var charsetLength = len(charset)

var cache map[string]bool

func generateRandomInt(min, max int) int {
	return random.Intn(max-min) + min
}

func generateRandomName() (name string) {
	for idx := 1; idx <= 2; idx++ {
		randomNumber := generateRandomInt(0, charsetLength-1)
		name += string(charset[randomNumber])
	}

	randomNumber := generateRandomInt(100, 999)
	name += strconv.Itoa(randomNumber)

	return name
}

// Name get the name of the robot
func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}

	var count int

	for {
		name := generateRandomName()

		count++

		if count > 1000000 {
			return "", errors.New("Retried name generation ten times and could not get unique name")
		}

		if cache[name] {
			continue
		}

		r.name = strings.ToUpper(name)
		break
	}

	return r.name, nil
}

// Reset changes the name of a robot to an empty string
func (r *Robot) Reset() {
	r.name = ""
}
