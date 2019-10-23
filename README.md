# HeyDon-Lee/utf8Util
* A Go Programming package for processing UTF-8 character data.

## Features
* high performance
* easy to use

## Install
```
go get github.com/HeyDon-Lee/utf8Util
```

## Examples
```go
package main

import "github.com/HeyDon-Lee/utf8Util"

func main() {
	// Gets the position of the string.
	println(utf8Util.Index(`你好！世界`, `世界`))
	println(utf8Util.Index(`こんにちは！世界`, `世界`))
	println(utf8Util.Index(`안녕하세요！세계`, `세계`))

	// Gets the position of the character.
	println(utf8Util.IndexOne(`你好！世界`, `世`))
	println(utf8Util.IndexOne(`こんにちは！世界`, `世`))
	println(utf8Util.IndexOne(`안녕하세요！세계`, `세`))

	// Gets the last position of the string.
	println(utf8Util.LastIndex(`你好！世界！世界`, `世界`))
	println(utf8Util.LastIndex(`こんにちは！世界！世界`, `世界`))
	println(utf8Util.LastIndex(`안녕하세요！세계！세계`, `세계`))

	// Gets the last position of the character.
	println(utf8Util.LastIndexOne(`你好！世界！世界`, `世`))
	println(utf8Util.LastIndexOne(`こんにちは！世界！世界`, `世`))
	println(utf8Util.LastIndexOne(`안녕하세요！세계！세계`, `세`))

	// Gets the substring, gets it by starting the index and the length of the substring, and returns the substring.
	println(utf8Util.SubStr(`你好！世界`, 3, 2))
	println(utf8Util.SubStr(`こんにちは！世界！世界`, 3, 2))
	println(utf8Util.SubStr(`안녕하세요！세계！세계`, 3, 2))

	// Gets the substring, gets it by starting and closing the index, and returns the substring.
	println(utf8Util.SubString(`你好！世界`, 3, 4))
	println(utf8Util.SubString(`こんにちは！世界！世界`, 3, 4))
	println(utf8Util.SubString(`안녕하세요！세계！세계`, 3, 4))

	// The location of the fetched substring of large text.
	println(utf8Util.LargeIndex(`你好！世界`, `世界`))
	println(utf8Util.LargeIndex(`こんにちは！世界`, `世界`))
	println(utf8Util.LargeIndex(`안녕하세요！세계`, `세계`))
}
```
