package main

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"
)

func creatWord(minLen, maxLen int) string {
	str := ""
	length := rand.Intn(maxLen - minLen + 1) + minLen
	for i := 0; i < length; i++ {
		str += string(int('a') + rand.Intn(26))
	}
	str += "\r\n"
	return str
}

func TestCreateWord(t *testing.T)  {
	rand.Seed(time.Now().UnixNano())
	word := creatWord(1, 100)
	fmt.Println(word, len(word))
}

func createLargeFile(fileSize int) {
	file, err := os.OpenFile("./LargeFile.txt", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	curSize := 0
	for curSize < fileSize {
		word := creatWord(1, 100)
		n, err := file.Write([]byte(word))
		if err != nil {
			panic(err)
		}
		curSize += n
	}
}

func TestCreateLargeFile(t *testing.T) {
	createLargeFile(100 * (1 << 30))
}