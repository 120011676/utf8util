# hdUnicode
Golang gets the position of the Unicode character.

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
	println(hdUnicode.UniIndex(`안녕하세요! 세계`, `세계`))
	println(hdUnicode.UniIndex(`你好世界`, `世界`))
	println(hdUnicode.UniIndex(`你好世界`, `世界`))
	println(hdUnicode.UniIndex(`你好世界`, `世界`))
}
```
