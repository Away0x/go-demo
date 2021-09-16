```bash
# 转换单词
go run . word -m=1 -s=hello
# 获取时间
go run . time now
go run . time calc -c="2029-09-04 12:02: 33" -d=5m
# sql to struct
go run . sql struct --username=root --password=a12345678 --db=dbname --table=tablename
# json to struct
go run . json struct -s='{"a": 123, "b": "abc", "c": true, "d": false, "e": [1, "2"]}'
```