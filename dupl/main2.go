package main

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"fmt"
)

var subslices [][]byte
var newArray [][]byte

func main() {
	var out int
	if len(os.Args) < 2 {
		log.Print("Command-line argument not found!\n")
		os.Exit(2)
	}
	log.Println("start")
	inputFile := os.Args[1]
	outputFile := inputFile + ".new"
	dat, err := ioutil.ReadFile(inputFile)
	check(err)
	subslices = bytes.Split(bytes.Replace(dat, []byte("\r\n"), []byte("\n"), -1), []byte("\n"))
	dat = nil
	in := len(subslices)
	procenti := 0.0
	procenti2 := 0.0
	for i := range subslices {
		procenti = (float64(i) / float64(in)) * 100
		if procenti >= procenti2 + 10 {
			procenti2 = procenti
			fmt.Printf("%d%%.. ",int(procenti))
		}
		if len(subslices[i]) == 0 {
			continue
		}
		if !checkDupl(subslices[i]) {
			newArray = append(newArray, subslices[i])
		}
	}
	fmt.Print("100%..\n")
	for ix := range newArray {
		dat = append(dat, append(newArray[ix], "\n"...)...)
		out++
	}
	subslices = nil
	err = ioutil.WriteFile(outputFile, dat, 0644)
	check(err)
	log.Println("finish")
	log.Print("Input rows: ", in, " Output rows: ", out, "\n", "Press 'Enter' to exit:")
	bufio.NewReader(os.Stdin).ReadString('\n')
}

func checkDupl(in []byte) bool {
	for i := range newArray {
		if bytes.Equal(in, newArray[i]) {
			return true
		}
	}
	return false
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

