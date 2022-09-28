package util

import (
	"math/rand"
	"strings"
	"time"
)

var LARGE_CHARS = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
var SMALL_CHARS = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
var NUM_CHARS = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}

func init() {
	rand.Seed(time.Now().UnixNano())
}

/*
RandStringAll  生成指定长度随机字符串([a-zA-Z0-9])
*/
func RandStringAll(lenNum int) string {
	str := strings.Builder{}
	chars := append(append(LARGE_CHARS, SMALL_CHARS...), NUM_CHARS...)
	length := len(chars)
	for i := 0; i < lenNum; i++ {
		l := chars[rand.Intn(length)]
		str.WriteString(l)
	}
	return str.String()
}

/*
RandStringSmall  生成指定长度随机字符串([a-z])
*/
func RandStringSmall(lenNum int) string {
	str := strings.Builder{}
	length := len(SMALL_CHARS)
	for i := 0; i < lenNum; i++ {
		str.WriteString(SMALL_CHARS[rand.Intn(length)])
	}
	return str.String()
}

/*
RandStringLarge  生成指定长度随机字符串([A-Z])
*/
func RandStringLarge(lenNum int) string {
	str := strings.Builder{}
	length := len(LARGE_CHARS)
	for i := 0; i < lenNum; i++ {
		str.WriteString(LARGE_CHARS[rand.Intn(length)])
	}
	return str.String()
}

/*
RandStringNum  生成指定长度随机字符串([0-9])
*/
func RandStringNum(lenNum int) string {
	str := strings.Builder{}
	length := len(NUM_CHARS)
	for i := 0; i < lenNum; i++ {
		l := NUM_CHARS[rand.Intn(length)]
		str.WriteString(l)
	}
	return str.String()
}

/*
RandStringLargeSmall  生成指定长度随机字符串([a-zA-Z])
*/
func RandStringLargeSmall(lenNum int) string {
	str := strings.Builder{}
	chars := append(LARGE_CHARS, SMALL_CHARS...)
	length := len(chars)
	for i := 0; i < lenNum; i++ {
		l := chars[rand.Intn(length)]
		str.WriteString(l)
	}
	return str.String()
}

/*
RandStringLargeNum  生成指定长度随机字符串([A-Z0-9])
*/
func RandStringLargeNum(lenNum int) string {
	str := strings.Builder{}
	chars := append(LARGE_CHARS, NUM_CHARS...)
	length := len(chars)
	for i := 0; i < lenNum; i++ {
		l := chars[rand.Intn(length)]
		str.WriteString(l)
	}
	return str.String()
}

/*
RandStringSmallNum  生成指定长度随机字符串([a-z0-9])
*/
func RandStringSmallNum(lenNum int) string {
	str := strings.Builder{}
	chars := append(NUM_CHARS, SMALL_CHARS...)
	length := len(chars)
	for i := 0; i < lenNum; i++ {
		l := chars[rand.Intn(length)]
		str.WriteString(l)
	}
	return str.String()
}
