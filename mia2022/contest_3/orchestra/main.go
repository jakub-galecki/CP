package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var in = bufio.NewReader(os.Stdin)

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
func readStrings() []string   { return strings.Split(_scln(), " ") }
func readInt(s ...string) int { t, _ := strconv.Atoi(_sc(s)); return t }
func rints() []int {
	strs := readStrings()
	arr := make([]int, len(strs))
	for i, s := range strs {
		arr[i] = readInt(s)
	}
	return arr
}
func write(arg ...interface{}) { fmt.Print(arg...) }

func main() {
	args := rints()
	var coordinates [][]int
	for i := 0; i < args[2]; i++ {
		coordinates = append(coordinates, rints())
	}
	sum := 0
	for sx := 1; sx <= args[0]; sx++ {
		for sy := 1; sy <= args[1]; sy++ {
			for ex := sx; ex <= args[0]; ex++ {
				for ey := sy; ey <= args[1]; ey++ {
					tmp := 0
					for _, c := range coordinates {
						if (sx <= c[0] && c[0] <= ex) && (sy <= c[1] && c[1] <= ey) {
							tmp++
						}
					}
					if tmp >= args[3] {
						sum++
					}
				}
			}
		}
	}
	write(sum)
}
