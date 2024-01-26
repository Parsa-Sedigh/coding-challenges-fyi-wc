package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func printSeparator() {
	fmt.Println("-------------------------------------")
}

func goToBeginning(file *os.File) {
	// put the file read position back to the beginning, so the file can be read from the beginning, again
	if _, err := file.Seek(0, io.SeekStart); err != nil {
		log.Fatal(err)
	}
}

func main() {
	var options Options

	flag.BoolVar(&options.countBytes, "c", false, "count bytes")
	flag.BoolVar(&options.countLines, "l", false, "count lines")
	flag.BoolVar(&options.countWords, "w", false, "count words")
	flag.Parse()

	//if !options.countBytes && !options.countLines {
	//	options.countBytes = true
	//	options.countLines = true
	//}

	filename := flag.Arg(0)

	fmt.Printf("options: %+v\nfilename: %s\n", options, filename)

	if filename == "" {
		filename = "test.txt"
	}

	file, err := os.Open(filename)
	if err != nil {
		var pathError *os.PathError

		if errors.As(err, &pathError) {
			fmt.Println("file's path: ", pathError.Path)
		}

		log.Fatal(err)
	}

	defer file.Close()

	if options.countBytes {
		bytesNum := handleBytes(file)

		printSeparator()
		fmt.Println("number of bytes: ", bytesNum)
	}

	goToBeginning(file)

	if options.countLines {
		linesNum := handleLines(file)

		printSeparator()
		fmt.Println("Total number of lines: ", linesNum)
	}

	goToBeginning(file)

	if options.countWords {
		wordsNum := handleWords(file)

		printSeparator()
		fmt.Println("Total number of words: ", wordsNum)
	}
}

func handleBytes(file *os.File) int {
	bytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	return len(bytes)
}

func handleLines(file *os.File) int {
	scanner := bufio.NewScanner(file)
	var linesNum int

	fmt.Println("start reading ...")

	// while scanner.Scan() returns true, continue
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		linesNum++
		//fmt.Println(linesNum, "lines found", scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	///////////////// OR: /////////////////

	//for {
	//	success := scanner.Scan()
	//
	//	if !success {
	//		if err := scanner.Err(); err != nil {
	//			log.Fatal(err)
	//		} else {
	//			log.Println("Scan completed and reached EOF; done reading")
	//
	//			break
	//		}
	//	} else {
	//		linesNum++
	//		fmt.Println(linesNum, "lines found", scanner.Text())
	//	}
	//}

	return linesNum
}

func handleWords(file *os.File) int {
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	var wordsNum int

	for scanner.Scan() {
		wordsNum++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return wordsNum
}
