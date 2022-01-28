package robotname

import (
	"crypto/rand"
	"errors"
	"log"
	"math/big"
	"strconv"
	"strings"
)

// Define the Robot type here.
type Robot struct {
	name string
}

var charset = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var charsetLength = len(charset)
var maxNameLength = 5

var cache map[string]bool

func generateRandomInt(max int64) (int64, error) {
	nBig, err := rand.Int(rand.Reader, big.NewInt(max))
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return nBig.Int64(), nil
}

func generateRandomName() (name string, err error) {
	for idx := 1; idx <= maxNameLength; idx++ {
		if idx < 3 {
			random, err := generateRandomInt(int64(charsetLength))
			if err != nil {
				return "", err
			}
			name += string(charset[random])
		} else {
			random, err := generateRandomInt(int64(9))
			if err != nil {
				return "", err
			}
			name += strconv.Itoa(int(random))
		}
	}

	return name, err
}

func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}

	var count int

	for {
		name, err := generateRandomName()

		if err != nil {
			return "", err
		}

		count++

		if count > 1000 {
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

func (r *Robot) Reset() {
	r.name = ""
}
