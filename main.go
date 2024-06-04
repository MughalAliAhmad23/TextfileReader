package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"sync"
	"time"
)

func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}

func SpaceCounter(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	spaces := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			spaces++
		}
	}
	fmt.Println("totalspaces:", spaces)
}

func VowelsCounter(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	vowels := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u' {
			vowels++
		}

	}
	fmt.Println("total vowels:", vowels)
}

func LineCounter(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	lines := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '.' {
			lines++
		}
	}
	fmt.Println("total lines :", lines)
}

func Wordfrequeny(s string, wg *sync.WaitGroup) {
	defer wg.Done()

	wordcount := make(map[string]int)

	split := strings.Split(s, " ")
	for i := 0; i < len(split); i++ {
		if value, key := wordcount[split[i]]; key {
			wordcount[split[i]] = value + 1
		} else {
			wordcount[split[i]] = 1
		}
	}

	fmt.Println("word frequencies:", wordcount)
}

func Wordcounter(s string, wg *sync.WaitGroup) {

	defer wg.Done()

	spaces := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			spaces++ 
		}
	}
	totalwords := spaces + 1

	fmt.Println("total words:", totalwords)
}

func main() {

	var wg sync.WaitGroup
	wg.Add(5)

	defer timer("main")()

	filepath := "textfile.txt"

	file, ferr := os.Open(filepath)
	if ferr != nil {
		fmt.Println("Error in file opening", ferr)
		return
	}

	defer file.Close()

	filedata, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println("Error in reading the file", err)
		return
	}

	go Wordfrequeny(string(filedata), &wg)

	go SpaceCounter(string(filedata), &wg)

	go Wordcounter(string(filedata), &wg)

	go VowelsCounter(string(filedata), &wg)

	go LineCounter(string(filedata), &wg)

	wg.Wait()
	fmt.Println("main exists")
}
