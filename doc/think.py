'''
思路：
位图1，用于判断是否存在该词。对于每次读进来的词，计算哈希值，相应比特位置1。
位图2，用于标志是否重复。对于读进来的并且是被位图1标志过存在的词，则置1
队列，用于保存不重复词。队尾保留最新不重复词，每次push都是在队尾，pop则不一定
（改用队列，主要是为了防止大文件都是不相同词时，要遍历整个hashmap，相当于遍历大文件两遍）
'''


# 伪码
# 遍历文件
for word in largeFile:
    if bitmap1.isExist(word):
        bitmap2.add(word)
        pop word from dueue
    else:
        bitmap1.add(word)
        push word to dueue
        if len(dueue) > maxSize:    # 推算每次I/O文件的大小和队列、两个位图共16GB得 maxSize = 7GB
            write dueue to disk

# 结算结果
firstWord = dueue[0]    # 此时内存的第一个不重复词
# read data from disk
while word = read(disk):
    if bitmap2.isExist(word)
        continue
    else:
        break

if word:
    firstWord = word    # 如果硬盘有更早的第一个不重复的词，更新
