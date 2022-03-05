# MiNiTool
## 1. 单词转换
## 2. 翻译（转换）

### 1. 单词转换
```go
go run main.go word -h 查看帮助
--str[-s] 输入单词
--mode[-m] 模式
"1：全部转大写",
"2：全部转小写",
"3：下划线转大写驼峰",
"4：下划线转小写驼峰",
"5：驼峰转下划线",

go run main.go word -s=hotDog -m=1
输出结果：HOTDOG
```

### 2. 翻译 [支持翻译后转换]
```
go run main.go trans -h 查看帮助
--words[-w] 输入待翻译的词语
--from[-f] 源语言[default:auto]
--to[-t] 目标语言[default:en]
--translator[-T] 翻译器[default:百度翻译]
--mod[-m] 模式[default:不转换]
"1：全部转大写",
"2：全部转小写",
"3：下划线转大写驼峰",
"4：下划线转小写驼峰",
```