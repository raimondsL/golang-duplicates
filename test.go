package main

import (
	"io/ioutil"
	"math/rand"
	"time"
	"os"
	"strconv"
	"fmt"
)

func main() {
	var dat []byte
	if len(os.Args) < 2 {
		fmt.Print("Command-line argument not found!\n")
		os.Exit(2)
	}
	x, _ := strconv.Atoi(os.Args[1])
	for i := 0; i < x; i++ {
		dat = append(dat,append([]byte(RandStringRunes(16)),"\n"...)...)
	}
	ioutil.WriteFile("file.txt",dat,0644)
}

func init() {
    rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = letterRunes[rand.Intn(len(letterRunes))]
    }
    return string(b)
}