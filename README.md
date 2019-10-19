# HeyDon-Lee/hdUnicode
* Golang gets the position of the Unicode character.
* Golang 获取 Unicode 字符所在位置

# Install / 安装
```
go get github.com/HeyDon-Lee/hdUnicode
```

# Examples / 代码实例
```go
package main

import "github.com/HeyDon-Lee/hdUnicode"

func main() {
	// Get the position of the string
	println(hdUnicode.UniIndex(`你好！世界`, `世界`))
	println(hdUnicode.UniIndex(`こんにちは！世界`, `世界`))
	println(hdUnicode.UniIndex(`안녕하세요！세계`, `세계`))

	// Get the position of the character
	println(hdUnicode.UniIndexOne(`你好！世界`, `世`))
	println(hdUnicode.UniIndexOne(`こんにちは！世界`, `世`))
	println(hdUnicode.UniIndexOne(`안녕하세요！세계`, `세`))
}
```
