package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
	"bufio"
)

func CompareSize() error{
	fileStat1, fileStat2, err := StatFiles(FilePath1, FilePath2)

	if err != nil {
		return err
	}

	fmt.Println("\n### Comparison of file sizes ###")
	fmt.Printf("\nSize of 1st file: %d byte | Size of 2nd file: %d byte\n", fileStat1.Size(), fileStat2.Size())
	if fileStat1.Size() == fileStat2.Size() {
		fmt.Println("\nFiles are identical of size")
	} else {
		fmt.Println("\nFiles differ in size")
	}
	return nil
}

func CompareContent() error{

	file1, err := os.Open(FilePath1)
	if err != nil {
		log.Println("File can't be open", err)
		return err
	}
	defer file1.Close()

	file2, err := os.Open(FilePath2)
	if err != nil {
		log.Println("File can't be open", err)
		return err
	}
	defer file2.Close()

	reader1 := bufio.NewReader(file1)
    reader2 := bufio.NewReader(file2)
    
	fmt.Println("\n### Comparison of file contents ###")

    for {
        byte1, err1 := reader1.ReadByte()
        byte2, err2 := reader2.ReadByte()

        if err1 != nil && err1 != io.EOF {
			log.Println("Reading error", err1)
			return err1
        }
        if err2 != nil && err2 != io.EOF {
            log.Println("Reading error", err2)
			return err2
        }

        if err1 == io.EOF && err2 == io.EOF {
            break
        }

        if byte1 != byte2 {
			fmt.Println("\nFiles differ in content")
			return nil
        }
    }

    fmt.Println("\nFiles are identical in content")
    return nil
} 

func CompareHash() error{

	file1, err := os.Open(FilePath1)
	if err != nil {
		log.Println("File can't be open", err)
		return err
	}
	defer file1.Close()

	file2, err := os.Open(FilePath2)
	if err != nil {
		log.Println("File can't be open", err)
		return err
	}
	defer file2.Close()

	hash1 := sha256.New()
	hash2 := sha256.New()

	_, err1 := io.Copy(hash1, file1)
	if err1 != nil {
		log.Println("Copy error", err1)
		return err1
	}

	_, err2 := io.Copy(hash2, file2)
	if err2 != nil {
		log.Println("Copy error", err2)
		return err2
	}

	fileHash1 := hash1.Sum(nil)
	fileHash2 := hash2.Sum(nil)

	fmt.Println("\n### Comparison of file SHA-256 ###")
	fmt.Printf("\nHash of 1st file: %x\n", fileHash1)
	fmt.Printf("Hash of 2nd file: %x\n", fileHash2)

	for i := range fileHash1 {
		if fileHash1[i] != fileHash2[i] {
			fmt.Println("\nFiles differ in SHA-256")
			return nil
		}
	}
	fmt.Println("\nFiles are identical in SHA-256")
	return nil
}

func CompareAll(mode string){
    CompareSize()
    CompareContent()
    CompareHash()
}