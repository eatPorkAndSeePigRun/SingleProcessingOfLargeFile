package main

import (
	"fmt"
	"sync"
)

const smallFileNum = 42
const smallFileSize = 100 * (1 << 30) / smallFileNum
var waitGroup = sync.WaitGroup{}


func main() {
	// 初始化位图，哈希表
	bitmap := NewBitMap(1 << 29)
	var hashmap map[string]int64
	// indexOfWord用于标记字符串在大文件种出现的位置
	var indexOfWord int64
	// 一组用于布隆过滤器算法的哈希函数
	hashFunc := []func(string) int{DEKHash, APHash, JSHash, FNVHash}


	// 启动多个协程将大文件分割成多个小文件
	// 后每次I/O一个小文件，位图中若没被所有哈希函数算出哈希值相应的比特的标志，则当前元素只出现一次
	// 用哈希表维护只出现一次的元素，以及其在大文件中出现的位置
	HandleLargeFile("./LargeFile.txt")
	for i := 1; i <= smallFileNum; i++ {
		go ScanningSmallFile(i, bitmap, hashmap, &indexOfWord, hashFunc)
	}
	waitGroup.Add(smallFileNum)
	waitGroup.Wait()


	// 查询是否存在不重复的词，后输出
	if word, isFind := FindFirstNonRepetitiveWord(hashmap); isFind {
		fmt.Println("第一个不重复的词为： ", word)
	} else {
		fmt.Println("不存在不重复的词")
	}
}

