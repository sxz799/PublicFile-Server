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
RandAllString  生成指定长度随机字符串([a-zA-Z0-9])
*/
func RandAllString(lenNum int) string {
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
RandNumString  生成指定长度随机字符串([0-9])
*/
func RandNumString(lenNum int) string {
	str := strings.Builder{}
	length := len(NUM_CHARS)
	for i := 0; i < lenNum; i++ {
		l := NUM_CHARS[rand.Intn(length)]
		str.WriteString(l)
	}
	return str.String()
}

/*
RandString  生成指定长度随机字符串([a-zA-Z])
*/
func RandString(lenNum int) string {
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
RandSmallString  生成指定长度随机字符串([a-z])
*/
func RandSmallString(lenNum int) string {
	str := strings.Builder{}
	length := len(SMALL_CHARS)
	for i := 0; i < lenNum; i++ {
		str.WriteString(SMALL_CHARS[rand.Intn(length)])
	}
	return str.String()
}

/*
RandLargeString  生成指定长度随机字符串([A-Z])
*/
func RandLargeString(lenNum int) string {
	str := strings.Builder{}
	length := len(LARGE_CHARS)
	for i := 0; i < lenNum; i++ {
		str.WriteString(LARGE_CHARS[rand.Intn(length)])
	}
	return str.String()
}
