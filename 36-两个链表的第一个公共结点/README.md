# 第一种：暴力方法
### 思路理解
遍历两个链表，判断他们第一次相等值的位置，并进行返回



# 第二种：右对齐方法
### 思路理解
由于他们存在公共链表区域，所以一定是Y结构的，则可以通过先对齐的方式，这样可以在一个循环里完成所有步骤



# 第三种：hashmap方法
### 思路理解
根据其中一个链表创建hashmap，并逐一对比另一个链表


### 解题过程
第一步：先创建一个key为链表的指针，value为布尔类型的hashmap，将链表存储为key，并默认value为ture

第二步：for循环遍历另外一个链表，并再循环里将该链表作为hashmap的key判断是否ok，找到改公共节点
