package main

import "flag"

var FilePath1, FilePath2, Mode string

func main() {

	//home/darksin/Dev/diffChecker/testRead1.txt
	//../diffChecker/testRead2.txt

	flag.StringVar(&FilePath1, "file1", "", "The path to first file")
	flag.StringVar(&FilePath2, "file2", "", "The path to second file")
	flag.StringVar(&Mode, "mode", "", "Selecting a file comparison option")

	flag.Parse()

	switch Mode {
	case "size":
		CompareSize()
	case "content":
		CompareContent() 
	case "hash":
		CompareHash()
	case "all":
        CompareSize()
		CompareContent()
        CompareHash() 
	}

	
}
