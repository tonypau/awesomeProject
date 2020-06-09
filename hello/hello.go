//package main
//
//import "fmt"
//
//const englishHelloPrefix = "Hello, "
//
//func Hello(name string) string {
//	// if name == "" {
//	// 	name = "World"
//	// }
//	return englishHelloPrefix + name
//}
//
//func main() {
//	fmt.Println(Hello(""))
//}

package main

import (
	"fmt"
	"strings"
)

const (
	alphabetStr = "1JedR8LNFY2j6MrhkBSADUyfP5amuH9xQCX4VqbgpsGtnW7vc3TwKEo"
	blockSize   = 22
	minLength   = 0
)

type UrlEncoder struct {
	Mask     int64
	Mapping  []uint
	Alphabet []string
}

func NewUrlEncoder() *UrlEncoder {
	var mask int64
	var mapping []uint
	var alphabet []string

	mask = 1<<blockSize - 1
	for i := blockSize - 1; i >= 0; i-- {
		mapping = append(mapping, uint(i))
	}
	for _, char := range alphabetStr {
		alphabet = append(alphabet, string(char))
	}

	return &UrlEncoder{
		Mask:     mask,
		Mapping:  mapping,
		Alphabet: alphabet,
	}
}

func (encoder *UrlEncoder) EncodeUrl(num int64) string {
	return encoder.enbaseOp(encoder.encodeOp(num))
}

func (encoder *UrlEncoder) enbaseOp(num int64) string {
	var padding string
	result := encoder.enbase(num)
	delta := minLength - len(result)
	fmt.Println(0 + delta)
	if delta > 0 {
		padding = strings.Repeat(encoder.Alphabet[0], delta)
	} else {
		padding = ""
	}
	return fmt.Sprintf("%s%s", padding, result)
}

func (encoder *UrlEncoder) enbase(num int64) string {
	n := int64(len(encoder.Alphabet))
	if num < n {
		return encoder.Alphabet[num]
	} else {
		return fmt.Sprintf("%s%s", encoder.enbaseOp(num/n), encoder.Alphabet[num%n])
	}
}

func (encoder *UrlEncoder) encodeOp(num int64) int64 {
	return (num & ^encoder.Mask) | encoder.encode(num&encoder.Mask)
}

func (encoder *UrlEncoder) encode(num int64) int64 {
	var result int64 = 0
	for pos, val := range encoder.Mapping {
		if num&(1<<uint(pos)) > 0 {
			result = result | (1 << val)
		}
	}
	return result
}

func main() {
	fmt.Println("hello world")
	var encoder *UrlEncoder = NewUrlEncoder()
	fmt.Println(encoder.EncodeUrl(833))
}
