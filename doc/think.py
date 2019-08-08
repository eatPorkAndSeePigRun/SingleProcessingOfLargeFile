'''
思路：
从后往读大文件，保留最新的不重复单词

位图1，用于判断是否存在该词。对于每次读进来的词，计算哈希值，相应比特位置1。
位图2，用于标志是否重复。对于读进来的并且是被位图1标志过存在的词，则置1
队列，用于保存不重复词。队尾保留最新不重复词，每次push都是在队尾，pop则不一定
（改用队列和倒叙读大文件，主要是为了防止大文件都是不相同词时，要遍历整个hashmap，相当于遍历大文件两遍）
'''


# 伪码
# 遍历文件
for word in largeFile[::-1]:
    if bitmap1.isExist(word):
        bitmap2.add(word)
        pop word from dueue
    else:
        bitmap1.add(word)
        push word to dueue
        if len(dueue) > 4GB :
            write dueue to disk

# 结算结果
firstWord
if dueue:
    firstWord = dueue[-1]
else:
    # read data from disk
    while firstWord = read(disk):   # 也是倒着读
        if bitmap2.isExist(word)
            continue
        else:
            break
