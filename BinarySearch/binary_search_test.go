package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	ge "github.com/og/x/error"
	gtest "github.com/og/x/test"
	"io/ioutil"
	"os"
	"path"
	"sort"
	"testing"
)

var newLine = []byte("\n")
var comma = []byte(`,`)
func rank(key int, list []int)  int{
	lo := 0
	hi := len(list)-1
	for ;lo<=hi; {
		mid := lo + (hi - lo) / 2
		switch {
		case key < list[mid]:
			hi = mid - 1
		case key > list[mid]:
			lo = mid + 1
		default:
			return mid
		}
	}
	return -1
}
func TestBasic(t *testing.T) {
	as := gtest.NewAS(t)
	whiteList := ReadFileInts("res/data/tinyW.txt")
	data := ReadFileInts("res/data/tinyT.txt")
	sort.Ints(whiteList)
	var result []int
	for _, n := range data {
		if rank(n, whiteList) < 0 {
			result = append(result, n)
			fmt.Println(n)
		}
	}
	as.Equal(result, []int{50, 99, 13})
}
func ReadFileInts(filename string) (ints []int) {
	filename = path.Join(ge.String(os.Getwd()), "../", filename)
	data, err := ioutil.ReadFile(filename) ; check(err)
	data = bytes.TrimSuffix(data, []byte("\n"))
	jsonb := []byte("[")
	jsonb = append(jsonb, bytes.ReplaceAll(data, newLine,  comma)...)
	jsonb = append(jsonb, []byte("]")...)
	err = json.Unmarshal(jsonb, &ints) ; check(err)
	return
}
func check(err error) {
	if err != nil {panic(err)}
}