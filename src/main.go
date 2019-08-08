package main

import (
	"fmt"
)

var bitmap1 *BitMap
var bitmap2 *BitMap
var dueue []string

func main() {
	// 初始化位图
	bitmap1 = NewBitMap(1 << 29)
	bitmap2 = NewBitMap(1 << 29)
	// 一组用于布隆过滤器算法的哈希函数
	hashFunc := []func(string) int{DEKHash, APHash, JSHash, FNVHash}


	HandleLargeFile("./LargeFile.txt", hashFunc)


	// 查询是否存在不重复的词，后输出
	if word, isFind := FindFirstNonRepetitiveWord(hashFunc); isFind {
		fmt.Println("第一个不重复的词为： ", word)
	} else {
		fmt.Println("不存在不重复的词")
	}
}

