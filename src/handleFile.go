package main

import (
	"bufio"
	"os"
	"strconv"
)

// 分割大文件，并把所有产生小文件放入到当前目录下的smartFile文件夹中
func HandleLargeFile(filePath string) {
	largeFile, err := os.Open(filePath)
	defer largeFile.Close()
	if err != nil {
		panic(err)
	}

	// 分割成小文件
	deleteAllSmartFiles()
	buf := make([]byte, smallFileSize)
	for i := 1; i <= smallFileNum; i++ {
		_, err := largeFile.Seek(int64((i-1)*smallFileSize), 0)
		if err != nil {
			panic(err)
		}
		_, err = largeFile.Read(buf)
		if err != nil {
			panic(err)
		}

		smallFileI, err := os.OpenFile("./smartFile"+strconv.Itoa(int(i))+".txt", os.O_CREATE|os.O_WRONLY, os.ModePerm)
		if err != nil {
			panic(err)
		}
		_, err = smallFileI.Write(buf)
		if err != nil {
			panic(err)
		}
		err = smallFileI.Close()
		if err != nil {
			panic(err)
		}
	}
}

// 每次分割前，清空存放小文件的文件夹
func deleteAllSmartFiles() {
	err := os.RemoveAll("./smartFile")
	if err != nil {
		panic(err)
	}
}

// 处理小文件，维护位图和哈希表，位图用于表示元素是否被访问过，哈希表维护只出现一次的元素及其位置信息
func ScanningSmallFile(index int, bitmap *BitMap, hashmap map[string]int64, indexOfWord *int64, hashFunc []func(string) int) {
	smallFileI, err := os.Open("./smartFile" + strconv.Itoa(int(index)) + ".txt")
	defer smallFileI.Close()
	if err != nil {
		panic(err)
	}

	// 每次读取一个词，维护位图和哈希表
	scanner := bufio.NewScanner(smallFileI)
	for scanner.Scan() {
		(*indexOfWord)++
		word := scanner.Text()

		if wordIsRepeat(word, bitmap, hashFunc) {
			delete(hashmap, word)
		} else {
			setWordRepeat(word, bitmap, hashFunc)
			hashmap[word] = *indexOfWord
		}
	}
}

// 用布隆过滤器算法查询元素是否访问过
func wordIsRepeat(word string, bitmap *BitMap, hashFunc []func(string) int) bool {
	for i := 0; i < len(hashFunc); i++ {
		num := hashFunc[i](word)
		if !bitmap.IsExist(uint(num % bitmap.max)) {
			return false
		}
	}
	return true
}

// 维护位图，把访问过的元素的哈希函数组响应哈希值的比特位置1
func setWordRepeat(word string, bitmap *BitMap, hashFunc []func(string) int) {
	for i := 0; i < len(hashFunc); i++ {
		num := hashFunc[i](word)
		bitmap.Add(uint(num % bitmap.max))
	}
}

// 查询第一不重复的词
func FindFirstNonRepetitiveWord(hashmap map[string]int64) (string, bool) {
	const int64Max = int64(^uint(0) >> 1)
	word, index := "", int64Max
	for k, v := range hashmap {
		if v < index {
			word = k
		}
	}

	if word != "" {
		return word, true
	} else {
		return "", false
	}
}
