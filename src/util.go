package main

// 本文件由一组哈希函数构成，用于布隆过滤器

// DEXHash算法
func DEKHash(str string) int {
	hash := len(str)
	for i := range str {
		hash = ((hash << 5) ^ (hash >> 27)) ^ i
	}
	return hash & 0x7FFFFFFF
}

// APHash算法
func APHash(str string) int {
	hash := 0
	for i := 0; i < len(str); i++ {
		if i&1 == 0 {
			hash ^= (hash << 7) ^ int(str[i]) ^ (hash >> 3)
		} else {
			hash ^= ^((hash << 11) ^ int(str[i]) ^ (hash >> 5))
		}
	}
	return hash
}

// FNVHash算法
func FNVHash(str string) int {
	p := 16777619
	hash := 2166136261
	for i := range str {
		hash = (hash ^ i) * p
	}
	hash += hash << 13
	hash ^= hash >> 7
	hash += hash << 3
	hash ^= hash >> 17
	hash += hash << 5
	return hash
}

// JSHash算法
func JSHash(str string) int {
	hash := 1315423911
	for i := range str {
		hash ^= (hash << 5) + i + (hash >> 2)
	}
	return hash & 0x7FFFFFFF
}
