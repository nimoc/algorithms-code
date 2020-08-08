package main

import (
	"fmt"
	ge "github.com/og/x/error"
	gtest "github.com/og/x/test"
	"io/ioutil"
	"os"
	"path"
	"sort"
	"strconv"
	"testing"
)

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
func rank2(key int, list []int)  int{
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
func TestBasic1(t *testing.T) {
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
func TestBasic2(t *testing.T) {
	as := gtest.NewAS(t)
	whiteList := ReadFileInts("res/data/tinyW.txt")
	data := ReadFileInts("res/data/tinyT.txt")
	sort.Ints(whiteList)
	var result []int
	for _, n := range data {
		if rank2(n, whiteList) < 0 {
			result = append(result, n)
			fmt.Println(n)
		}
	}
	as.Equal(result, []int{50, 99, 13})
}

// 读取文件并将内容转换为 []int
func ReadFileInts(filename string) (ints []int) {
	filename = path.Join(ge.String(os.Getwd()), "../", filename)
	data, err := ioutil.ReadFile(filename) ; check(err)
	newLine := byte(10)
	var temp []byte
	for _, b := range data {
		if b == newLine {
			if len(temp) == 0 {continue}
			i, err := strconv.Atoi(string(temp)) ; check(err)
			ints = append(ints, i)
			temp = []byte{}
			continue
		}
		temp = append(temp, b)
	}
	return
}
func check(err error) {
	if err != nil {panic(err)}
}
