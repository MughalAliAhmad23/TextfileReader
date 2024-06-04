package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}

func SpaceCounter(s string) int {
	spaces := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			spaces++
		}
	}
	return spaces
}

func VowelsCounter(s string) int {
	vowels := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u' {
			vowels++
		}

	}
	return vowels
}

func LineCounter(s string) int {
	lines := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '.' {
			lines++
		}
	}
	return lines
}

func Wordfrequeny(s string) map[string]int {
	var temp string

	wordcount := make(map[string]int)

	for i := 0; i < len(s); i++ {
		if s[i] == ' ' || i == len(s)-1 {
			if value, key := wordcount[temp]; key {
				wordcount[temp] = value + 1
			} else {
				wordcount[temp] = 1
			}
			temp = ""
		} else {
			temp += string(s[i])
		}
	}
	return wordcount
}

func Wordcounter(s string) int {
	spaces := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			spaces++
		}
	}
	totalwords := spaces + 1
	return totalwords
}

func main() {

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

	totalspaces := SpaceCounter(string(filedata))

	fmt.Println("total spaces :", totalspaces)

	totalwords := Wordcounter(string(filedata))

	fmt.Println("total words :", totalwords)

	totalvowels := VowelsCounter(string(filedata))

	fmt.Println("total vowels:", totalvowels)

	wordfrequencies := Wordfrequeny(string(filedata))

	fmt.Println(wordfrequencies)

	totallines := LineCounter(string(filedata))

	fmt.Println("total no of lines in paragraph", totallines)
}
