for range map是无序的，如果第一次循环到A，counter的值等于3，否则等于2
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
