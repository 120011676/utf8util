# HeyDon-Lee/utf8util
* A Go Programming package for processing UTF-8 character data.

## Install
```
go get github.com/HeyDon-Lee/utf8util
```

## Examples
```go
package main

import "github.com/HeyDon-Lee/utf8util"

func main() {
	// Gets the position of the string.
	println(utf8util.Index(`你好！世界`, `世界`))
	println(utf8util.Index(`こんにちは！世界`, `世界`))
	println(utf8util.Index(`안녕하세요！세계`, `세계`))

	// Gets the position of the character.
	println(utf8util.IndexOne(`你好！世界`, `世`))
	println(utf8util.IndexOne(`こんにちは！世界`, `世`))
	println(utf8util.IndexOne(`안녕하세요！세계`, `세`))

	// Gets the last position of the string.
	println(utf8util.LastIndex(`你好！世界！世界`, `世界`))
	println(utf8util.LastIndex(`こんにちは！世界！世界`, `世界`))
	println(utf8util.LastIndex(`안녕하세요！세계！세계`, `세계`))

	// Gets the last position of the character.
	println(utf8util.LastIndexOne(`你好！世界！世界`, `世`))
	println(utf8util.LastIndexOne(`こんにちは！世界！世界`, `世`))
	println(utf8util.LastIndexOne(`안녕하세요！세계！세계`, `세`))

	// Gets the substring, gets it by starting the index and the length of the substring, and returns the substring.
	println(utf8util.SubStr(`你好！世界`, 3, 2))
	println(utf8util.SubStr(`こんにちは！世界！世界`, 3, 2))
	println(utf8util.SubStr(`안녕하세요！세계！세계`, 3, 2))

	// Gets the substring, gets it by starting and closing the index, and returns the substring.
	println(utf8util.SubString(`你好！世界`, 3, 4))
	println(utf8util.SubString(`こんにちは！世界！世界`, 3, 4))
	println(utf8util.SubString(`안녕하세요！세계！세계`, 3, 4))

	// The location of the fetched substring of large text.
	println(utf8util.LargeIndex(`你好！世界`, `世界`))
	println(utf8util.LargeIndex(`こんにちは！世界`, `世界`))
	println(utf8util.LargeIndex(`안녕하세요！세계`, `세계`))
}
```
