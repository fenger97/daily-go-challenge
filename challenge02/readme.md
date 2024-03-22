counter的值可能是2，也可能是3
输出内容如下：
```bash
❯ go run main.go
C 23
B 22
2
❯ go run main.go
A 21
B 22
C 23
3
```
for range map是无序的，如果第一次循环到A，则输出3，否则输出2