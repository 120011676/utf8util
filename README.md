# HeyDon-Lee/hdUnicode
* A Go Programming package for processing UniCode character data.

## Features
* high performance
* easy to use

## Install
```
go get github.com/HeyDon-Lee/hdUnicode
```

## Examples
```go
package main

import "github.com/HeyDon-Lee/hdUnicode"

func main() {
	// Gets the position of the string.
	println(hdUnicode.Index(`你好！世界`, `世界`))
	println(hdUnicode.Index(`こんにちは！世界`, `世界`))
	println(hdUnicode.Index(`안녕하세요！세계`, `세계`))

	// Gets the position of the character.
	println(hdUnicode.IndexOne(`你好！世界`, `世`))
	println(hdUnicode.IndexOne(`こんにちは！世界`, `世`))
	println(hdUnicode.IndexOne(`안녕하세요！세계`, `세`))

	// Gets the last position of the string.
	println(hdUnicode.LastIndex(`你好！世界！世界`, `世界`))
	println(hdUnicode.LastIndex(`こんにちは！世界！世界`, `世界`))
	println(hdUnicode.LastIndex(`안녕하세요！세계！세계`, `세계`))

	// Gets the last position of the character.
	println(hdUnicode.LastIndexOne(`你好！世界！世界`, `世`))
	println(hdUnicode.LastIndexOne(`こんにちは！世界！世界`, `世`))
	println(hdUnicode.LastIndexOne(`안녕하세요！세계！세계`, `세`))

	// Gets the substring, gets it by starting the index and the length of the substring, and returns the substring.
	println(hdUnicode.SubStr(`你好！世界`, 3, 2))
	println(hdUnicode.SubStr(`こんにちは！世界！世界`, 3, 2))
	println(hdUnicode.SubStr(`안녕하세요！세계！세계`, 3, 2))

	// Gets the substring, gets it by starting and closing the index, and returns the substring.
	println(hdUnicode.SubString(`你好！世界`, 3, 4))
	println(hdUnicode.SubString(`こんにちは！世界！世界`, 3, 4))
	println(hdUnicode.SubString(`안녕하세요！세계！세계`, 3, 4))

	// The location of the fetched substring of large text.
	println(hdUnicode.LargeIndex(`你好！世界`, `世界`))
	println(hdUnicode.LargeIndex(`こんにちは！世界`, `世界`))
	println(hdUnicode.LargeIndex(`안녕하세요！세계`, `세계`))
}
```
