package dictionary

import (
	"time"
	"math/rand"
	"bufio"
	"os"
)

var words = make([]string, 0, 50)

func Load(fileName string) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return err
	} 
	return nil
}

func PickWord() string {
	rand.Seed(time.Now().Unix())
	i := rand.Intn(len(words))
	return words[i]
}