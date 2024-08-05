package basex

import (
	"math"
	"strings"
)

var l int
var strB string

func MustBaseInit(s string) {
	strB = s
	l = len(strB)
	if l == 0 {
		panic("Base String Need Init")
	}
}

// IntToString 十进制数转为62进制
func IntToString(s uint64) string {
	if s == 0 {
		return string(strB[0])
	}

	var seq []byte

	for s > 0 {
		mod := int(s) % l
		div := int(s) / l
		seq = append(seq, strB[mod])
		s = uint64(div)
		// 这样得到的seq是反过来的，需要翻转
	}
	reverse(seq)
	return string(seq)
}

func StringToInt(s string) uint64 {

	// 先翻转
	bl := []byte(s)
	temp := reverse(bl)

	var res float64
	for idx, v := range temp {
		base := math.Pow(float64(l), float64(idx))
		res += base * float64(strings.Index(strB, string(v)))
	}
	return uint64(res)
}

// reverse 翻转
func reverse(s []byte) []byte {
	for i, j := 0, len(s)-1; i < len(s)/2; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
