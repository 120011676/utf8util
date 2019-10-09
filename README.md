# HeyDon-Lee/hdUnicode
* Golang gets the position of the Unicode character.
* Golang 获取 Unicode 字符所在位置

# Install
```
go get github.com/HeyDon-Lee/hdUnicode
```

# Examples
```
package main

import "github.com/HeyDon-Lee/hdUnicode"

func main() {
	println(hdUnicode.UniIndex(`你好！世界`, `世界`))
	println(hdUnicode.UniIndex(`こんにちは！世界`, `世界`))
	println(hdUnicode.UniIndex(`안녕하세요！세계`, `세계`))
}
```
