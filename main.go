package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func printSeparator() {
	fmt.Println("-------------------------------------")
}

func main() {
	var options Options

	flag.BoolVar(&options.countBytes, "c", false, "count bytes")
	flag.BoolVar(&options.countLines, "l", false, "count lines")
	flag.Parse()

	//if !options.countBytes && !options.countLines {
	//	options.countBytes = true
	//	options.countLines = true
	//}

	filename := flag.Arg(0)

	fmt.Printf("options: %+v\n", options)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	if options.countBytes {
		bytes, err := io.ReadAll(file)
		if err != nil {
			log.Fatal(err)
		}

		printSeparator()
		fmt.Println("number of bytes: ", len(bytes))
	}

	if options.countLines {
		scanner := bufio.NewScanner(file)
		var linesNum uint

		fmt.Println("start reading ...")

		// while scanner.Scan() returns true, continue
		for scanner.Scan() {
			if err := scanner.Err(); err != nil {
				log.Fatal(err)
			}

			linesNum++
			fmt.Println(linesNum, "lines found", scanner.Text())
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

		printSeparator()
		fmt.Println("Total number of lines: ", linesNum)
	}
}
