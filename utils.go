package main

import (
	"fmt"
	"log"
	"os"
)

func StatFiles(filePath1, filePath2 string) (fileStat1, fileStat2 os.FileInfo){

    fileStat1,err1 := os.Stat(filePath1)
    if err1 != nil {
	    fmt.Println("File is not exist")
	    log.Fatal(err1)
    }

    fileStat2,err2 := os.Stat(filePath2)
    if err2 != nil {
	    fmt.Println("File is not exist")
	    log.Fatal(err2)
    }
	return fileStat1, fileStat2
}

