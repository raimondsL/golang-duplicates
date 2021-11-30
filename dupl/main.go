package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Print("Command-line argument not found!\n")
		os.Exit(2)
	}
	// log.Println("*********************")
	// log.Println("*                   *")
	// log.Println("* Delete duplicates *")
	// log.Println("* Keeping the order *")
	// log.Println("*                   *")
	// log.Println("*********************")
	inputFile := os.Args[1]
	outputFile := inputFile + ".new"
	dat, err := ioutil.ReadFile(inputFile) // []byte
	check(err)
	dat = bytes.Replace(dat, []byte("\r\n"), []byte("\n"), -1)
	dat = bytes.Trim(dat, "\n ") // trim nl and ws
	subslices := bytes.Split(dat, []byte("\n")) // [][]byte
	dat = nil
	subslices = MergeSort(subslices)
	l := len(subslices) - 2 // ain't stupid
	d := 0
	dat = append(dat, append(subslices[0], "\n"...)...)
	for i := 0; i < l; i++ {
		if bytes.Compare(subslices[i], subslices[i+1]) != 0 {
			dat = append(dat, append(subslices[i+1], "\n"...)...)
		} else {
			d++
		}
	}
	if d > 0 {
		err = ioutil.WriteFile(outputFile, dat, 0644)
		check(err)
	}
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

func Subtract(a [][]byte, d [][]byte) [][]byte {
	for i := range d {
		for j := range a {
			if bytes.Compare(a[j], d[i]) == 0 {
				a[j] = nil
			}
		}
	}
	return a
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
