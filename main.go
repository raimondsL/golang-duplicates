package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var dat []byte
var err error

func main() {
	if len(os.Args) < 2 {
		log.Print("Command-line argument not found!\n")
		os.Exit(2)
	}
	log.Println("*********************")
	log.Println("*                   *")
	log.Println("* Delete duplicates *")
	log.Println("* Keep the order... *")
	log.Println("*                   *")
	log.Println("*********************")
	inputFile := os.Args[1]
	outputFile := inputFile + ".new"
	dat, err = ioutil.ReadFile(inputFile) // []byte
	check(err)
	dat = bytes.Replace(dat, []byte("\r\n"), []byte("\n"), -1)
	dat = bytes.Trim(dat, "\n ") //trim existing or non existing \n and space
	dat = append(dat, "\n"...) //add one last \n
	var subslices [][]byte
	subslices = bytes.Split(dat, []byte("\n")) // [][]byte
	subslices = MergeSort(subslices)
	l := len(subslices) - 2 // ain't stupid
	d := 0
	for i := 0; i < l; i++ {
		if bytes.Compare(subslices[i], subslices[i+1]) == 0 {
			d++
			SubtractDat(subslices[i])
		}
	}
	err = ioutil.WriteFile(outputFile, dat, 0644)
	check(err)
	log.Println("finish")
	log.Println("Input rows: ", l+2, " Duplicate rows: ", d)
	fmt.Println("Press 'Enter' to exit:")
	bufio.NewReader(os.Stdin).ReadString('\n')
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func SubtractDat(s []byte) {
	s = append(s, "\n"...)
	l := len(s)
	idx := bytes.LastIndex(dat, s)
	if idx != -1 {
		fmt.Printf("from index %v len %v subtract %s",idx,l,s)
	}
}

func MergeSort(slice [][]byte) [][]byte {

	if len(slice) < 2 {
		return slice
	}

	mid := (len(slice)) / 2
	return Merge(MergeSort(slice[:mid]), MergeSort(slice[mid:]))
}

func Merge(left, right [][]byte) [][]byte {

	size, i, j := len(left)+len(right), 0, 0
	slice := make([][]byte, size, size)

	for k := 0; k < size; k++ {
		if i > len(left)-1 && j <= len(right)-1 {
			slice[k] = right[j]
			j++
		} else if j > len(right)-1 && i <= len(left)-1 {
			slice[k] = left[i]
			i++
		} else if bytes.Compare(left[i], right[j]) < 0 {
			slice[k] = left[i]
			i++
		} else {
			slice[k] = right[j]
			j++
		}
	}
	return slice
}
