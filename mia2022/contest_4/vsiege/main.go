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
func write(arg ...interface{}) { fmt.Println(arg...) }

func prfxsum(health []int, target []int64) {
	target[1] = int64(health[1])
	for i := 2; i < len(health); i++ {
		target[i] = target[i-1] + int64(warriors[i])
	}
}

func bsearch(cursor, damage, nworriros int, arrows *int64) int {
	if *arrows+int64(damage) < int64(warriors[cursor]) {
		*arrows += int64(damage)
		return cursor
	}
	a := cursor
	b := nworriros
	var killed int
	for a <= b {
		mid := int((a + b) / 2)
		if (prefix[mid] - prefix[cursor-1] - int64(*arrows)) <= int64(damage) {
			a = mid + 1
			killed = mid
		} else {
			b = mid - 1
		}
	}

	if killed == nworriros {
		*arrows = 0
		return 1
	}
	killed += 1
	hp_taken := prefix[killed] - prefix[cursor-1] - int64(*arrows) - int64(damage)
	*arrows = int64(warriors[killed]) - hp_taken
	return killed
}

var args []int
var warriors []int
var k []int
var prefix []int64

func main() {
	args = rints()
	warriors = rints()
	warriors = append([]int{0}, warriors...)
	k = rints()
	prefix = make([]int64, args[0]+1)
	prfxsum(warriors, prefix)
	warrior := 1
	var arrows int64 = 0
	for tick := 0; tick < args[1]; tick++ {
		warrior = bsearch(warrior, k[tick], args[0], &arrows)
		write(args[0] - warrior + 1)
	}
}
