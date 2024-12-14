package main

import (
	"log"
	"os"
)

func StatFiles(filePath1, filePath2 string) (fileStat1, fileStat2 os.FileInfo, err error){

    fileStat1,err1 := os.Stat(filePath1)
    if err1 != nil {
	    log.Println("File is not exist", err1)
		return nil , nil, err
    }

    fileStat2,err2 := os.Stat(filePath2)
    if err2 != nil {
	    log.Println("File is not exist", err2)
		return nil , nil, err
    }
	return fileStat1, fileStat2, nil
}

