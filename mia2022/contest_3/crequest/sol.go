package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ~~~~~ template ~~~~~

var in = bufio.NewReader(os.Stdin)

type float = float32
type double = float64
type long = int64
type ulong = uint64

func _scln() string {
	ln, _ := in.ReadString('\n')
	return strings.Trim(ln, " \r\n")
}
func _sc(s []string) string {
	if len(s) == 0 {
		return _scln()
	}
	return s[0]
}

func readLine() string              { return _scln() }
func readString() string            { return _scln() }
func readStrings() []string         { return strings.Split(_scln(), " ") }
func readBool(s ...string) bool     { t, _ := strconv.ParseBool(_sc(s)); return t }
func readByte(s ...string) byte     { t, _ := strconv.ParseUint(_sc(s), 10, 8); return byte(t) }
func readDouble(s ...string) double { t, _ := strconv.ParseFloat(_sc(s), 64); return t }
func readFloat(s ...string) float   { t, _ := strconv.ParseFloat(_sc(s), 32); return float32(t) }
func readInt(s ...string) int       { t, _ := strconv.Atoi(_sc(s)); return t }
func readLong(s ...string) long     { t, _ := strconv.ParseInt(_sc(s), 10, 64); return t }
func readULong(s ...string) ulong   { t, _ := strconv.ParseUint(_sc(s), 10, 64); return t }

/* ~~~~~~~~~~ */

// Too much overhead with universal method using interface{}, wait for the generics,
// but for now — just don't look at the redundancy of the code after that commentary :)

func readBools() []bool {
	strs := readStrings()
	arr := make([]bool, len(strs))
	for i, s := range strs {
		arr[i] = readBool(s)
	}
	return arr
}
func readBytes() []byte {
	strs := readStrings()
	arr := make([]byte, len(strs))
	for i, s := range strs {
		arr[i] = readByte(s)
	}
	return arr
}
func readDoubles() []double {
	strs := readStrings()
	arr := make([]double, len(strs))
	for i, s := range strs {
		arr[i] = readDouble(s)
	}
	return arr
}
func readFloats() []float {
	strs := readStrings()
	arr := make([]float, len(strs))
	for i, s := range strs {
		arr[i] = readFloat(s)
	}
	return arr
}
func readInts() []int {
	strs := readStrings()
	arr := make([]int, len(strs))
	for i, s := range strs {
		arr[i] = readInt(s)
	}
	return arr
}
func readLongs() []long {
	strs := readStrings()
	arr := make([]long, len(strs))
	for i, s := range strs {
		arr[i] = readLong(s)
	}
	return arr
}
func readULongs() []ulong {
	strs := readStrings()
	arr := make([]ulong, len(strs))
	for i, s := range strs {
		arr[i] = readULong(s)
	}
	return arr
}

func write(arg ...interface{})                 { fmt.Print(arg...) }
func writeLine(arg ...interface{})             { fmt.Println(arg...) }
func writeFormat(f string, arg ...interface{}) { fmt.Printf(f, arg...) }

// ~~~~ end template ~~~~

func reverse(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}

func main() {
	input := readInts()
	sum := 0

	for i := 1; i <= input[0]; i++ {
		palindrome_str := strconv.Itoa(i) + reverse(strconv.Itoa(i))
		palindrome_int, _ := strconv.Atoi(palindrome_str)
		sum = (sum + palindrome_int) % input[1]
	}

	write(sum)
}
