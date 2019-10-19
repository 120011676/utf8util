# HeyDon-Lee/hdUnicode
---
* A Go Programming package for processing UniCode character data.

## Features
---
* high performance
* easy to use

## Install
---
```
go get github.com/HeyDon-Lee/hdUnicode
```

## Examples
---
```go
package main

import "github.com/HeyDon-Lee/hdUnicode"

func main() {
	// Gets the position of the string.
	println(hdUnicode.UniIndex(`你好！世界`, `世界`))
	println(hdUnicode.UniIndex(`こんにちは！世界`, `世界`))
	println(hdUnicode.UniIndex(`안녕하세요！세계`, `세계`))

	// Gets the position of the character.
	println(hdUnicode.UniIndexOne(`你好！世界`, `世`))
	println(hdUnicode.UniIndexOne(`こんにちは！世界`, `世`))
	println(hdUnicode.UniIndexOne(`안녕하세요！세계`, `세`))

	// Gets the last position of the string.
	println(hdUnicode.UniLastIndex(`你好！世界！世界`, `世界`))
	println(hdUnicode.UniLastIndex(`こんにちは！世界！世界`, `世界`))
	println(hdUnicode.UniLastIndex(`안녕하세요！세계！세계`, `세계`))

	// Gets the last position of the character.
	println(hdUnicode.UniLastIndexOne(`你好！世界！世界`, `世`))
	println(hdUnicode.UniLastIndexOne(`こんにちは！世界！世界`, `世`))
	println(hdUnicode.UniLastIndexOne(`안녕하세요！세계！세계`, `세`))
}
```
