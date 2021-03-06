# 数据结构与算法

从工程角度理解和分析数据结构相关知识

## 目录

- [位图](https://github.com/chenhg5/CS-Interview/blob/master/dsa/bitmap.md)
- [哈希表](https://github.com/chenhg5/CS-Interview/blob/master/dsa/hash.md)
- [前缀树](https://github.com/chenhg5/CS-Interview/blob/master/dsa/trie.md)
- [leetcode](https://github.com/chenhg5/CS-Interview/blob/master/dsa/trie.md)
- [不可变数据结构](https://github.com/chenhg5/CS-Interview/blob/master/dsa/immutable.md)

## 哲学

- 离开数据量和数据访问模式来谈数据结构都是耍流氓

## 常量

熟悉各种常量可以帮助判断dsa的边界与瓶颈所在

### 数据类型字节数

- 1 bytes(字节) = 8 bit
- bytes, kb, mb, gb, tb 之间换算比例都是 1024

数据类型 | 说明	| 32位字节数 | 64位字节数 |	取值范围
-|-|-|-|-
bool |	布尔型	|1	|1|	true，false
char|	字符型	|1|	1|	-128~127
unsigned char|	无符号字符型	|1|	1|	0~255
short	|短整型|	2	|2	|-32768~32767
unsigned short	|无符号短整型|	2	|2|	0~65535
int|	整型|	4	|4	|-2147483648~2147483647
unsigned int|	无符号整型	|4	|4	|0~4294967295
long	|长整型|	4	|8	|–
unsigned long|	无符号长整型|	4	|8|	–
long long	|长整型|	8|	8|	-2^64~2^64-1
float|	单精度浮点数|	4|	4|	范围-2^128-2^128，精度为6~7位有效数字
double	|双精度浮点数|	8|	8|	范围-2^1024-2^1024，精度为15~16位
long double|	扩展精度浮点数|	8|	8|	范围-2^1024-2^1024，精度为15~16位

**一亿个整数**：1,0000,0000*4 Bytes = 381.4 MB = 0.3 GB

**一亿个字符串**（假设平均20个char）：1,0000,0000*20 Bytes = 1907 MB = 1.5 GB

### cpu计算速度

FLOP/s，Floating-point operations per second，每秒峰值速度，

- 一个 MFLOPS（megaFLOPS）等於每秒一佰万（=10^6）次的浮点运算，
- 一个 GFLOPS（gigaFLOPS）等於每秒拾亿（=10^9）次的浮点运算，
- 一个 TFLOPS（teraFLOPS）等於每秒万亿（=10^12）次的浮点运算，
- 一个 PFLOPS（petaFLOPS）等於每秒千万亿（=10^15）次的浮点运算，
- 一个 EFLOPS（exaFLOPS）等於每秒百亿亿（=10^18）次的浮点运算。

Intel 的 i5-2520M @2.5 Ghz 的处理器为例，其计算能力 2.5 * 4（4核）= 10 GFLOPS

### 内存

### io速度

- 固态硬盘
- 机械硬盘
- 内存
- cpu高速缓存
