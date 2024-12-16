package main

import "flag"

func main() {

	//home/darksin/Dev/diffChecker/testRead1.txt
	//../diffChecker/testRead2.txt

	var filePath1, filePath2, mode string

	flag.StringVar(&filePath1, "file1", "", "The path to first file")
	flag.StringVar(&filePath2, "file2", "", "The path to second file")
	flag.StringVar(&mode, "mode", "", "Selecting a file comparison option")

	flag.Parse()

	switch mode {
	case "size":
		CompareSize(filePath1, filePath2)
	case "content":
		CompareContent(filePath1, filePath2)
	case "hash":
		CompareHash(filePath1, filePath2)
	case "all":
		CompareAll(filePath1, filePath2)
	}
}