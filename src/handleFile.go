package main

import (
	"bufio"
	"os"
)

const readFileNum = 14
const bufSize = 100 * (1 << 30) / readFileNum

// 分次I/O读取大文件，并对每次读取的词调用相关函数维护位图和队列
func HandleLargeFile(filePath string, hashFunc []func(string) int) {
	largeFile, err := os.Open(filePath)
	defer largeFile.Close()
	if err != nil {
		panic(err)
	}

	rd := bufio.NewReaderSize(largeFile, bufSize)
	for {
		word, err := rd.ReadString('\n')
		if err != nil {
			panic(err)
		}
		if len(word) == 0 {
			break
		}
		handleWord(word, hashFunc)
	}
}

// 处理每个词
// 对于位图bitmap1，实际上是一个布隆算法原理，它相当于一个hashset，存放着大文件中词
// 对于位图bitmap2，也是一个布隆算法原理，它也相当于一个hashset，存放大文件中重复出现的词
// 它进行的流程是：
// 		如果是第一次访问该词，则将bitmap1的相应比特位置1，加入队列，快爆内存时则将队列写到硬盘，清空队列内存
// 		如果是重复访问，则将bitmap2的相应比特位置1，从队列出队
func handleWord(word string, hashFunc []func(string) int) {
	if hasWord(bitmap1, word, hashFunc) {
		addWord(bitmap2, word, hashFunc)
		popFromDueue(word)
	} else {
		addWord(bitmap1, word, hashFunc)
		pushToDueue(word)
		if len(dueue) > bufSize/100 {
			writeToDisk()
		}
	}
}

// 将队列数据存储到内存
func writeToDisk() {
	f, _ := os.OpenFile("spare.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	defer f.Close()
	for _, word := range dueue {
		_, err := f.Write([]byte(word + "\r\n"))
		if err != nil {
			panic(err)
		}
	}
	dueue = dueue[0:0]
}

func pushToDueue(word string) {
	dueue = append(dueue, word)
}

func popFromDueue(word string) {
	for i := 0; i < len(dueue); {
		if dueue[i] == word {
			dueue = append(dueue[:i], dueue[i+1:]...)
		} else {
			i++
		}
	}
}

// 用布隆过滤器算法查询元素是否访问过
func hasWord(bitmap *BitMap, word string, hashFunc []func(string) int) bool {
	for i := 0; i < len(hashFunc); i++ {
		num := hashFunc[i](word)
		if !bitmap.IsExist(uint(num % bitmap.max)) {
			return false
		}
	}
	return true
}

// 维护布隆过滤器的位图，把访问过的元素的哈希函数组响应哈希值的比特位置1
func addWord(bitmap *BitMap, word string, hashFunc []func(string) int) {
	for i := 0; i < len(hashFunc); i++ {
		num := hashFunc[i](word)
		bitmap.Add(uint(num % bitmap.max))
	}
}

// 查询第一不重复的词
// 先找出内存中队列的最早不重复词和硬盘中的最早不重复词，再比较可能哪边更早
// 存在则返回词，不存在则返回空
func FindFirstNonRepetitiveWord(hashFunc []func(string) int) (string, bool) {
	word := ""
	// 当前内存中的最早不重复词
	if dueue != nil {
		word = dueue[0]
	}

	largeFile, err := os.Open("spare.txt")
	defer largeFile.Close()
	if err != nil {
		panic(err)
	}

	// 遍历硬盘
	rd := bufio.NewReader(largeFile)
	for {
		tmpWord, err := rd.ReadString('\n')
		if err != nil {
			panic(err)
		}
		if len(tmpWord) == 0 {
			break
		}

		if hasWord(bitmap2, tmpWord, hashFunc) {
			continue
		} else {
			word = tmpWord	// 找到硬盘中最早不重复词，则终止循环
			break
		}
	}

	// 返回结果
	if word != "" {
		return word, true
	} else {
		return "", false
	}
}
