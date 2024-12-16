package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
)

func CompareSize(filePath1, filePath2 string) (bool, error) {

	fileStat1, err := os.Stat(filePath1)
	if err != nil {
		log.Println("File1 is not exist", err)
		return false, err
	}

	fileStat2, err2 := os.Stat(filePath2)
	if err2 != nil {
		log.Println("File2 is not exist", err2)
		return false, err
	}

	fmt.Println("\n### Comparison of file sizes ###")
	fmt.Printf("\nSize of 1st file: %d byte | Size of 2nd file: %d byte\n", fileStat1.Size(), fileStat2.Size())
	if fileStat1.Size() == fileStat2.Size() {
		fmt.Println("\nFiles are identical of size")
	} else {
		fmt.Println("\nFiles differ in size")
		return false, nil
	}
	return true, nil
}

func CompareContent(filePath1, filePath2 string) (bool, error) {

	file1, err := os.Open(filePath1)
	if err != nil {
		log.Println("File1 can't be open", err)
		return false, err
	}
	defer file1.Close()

	file2, err := os.Open(filePath2)
	if err != nil {
		log.Println("File2 can't be open", err)
		return false, err
	}
	defer file2.Close()

	reader1 := bufio.NewReader(file1)
	reader2 := bufio.NewReader(file2)

	fmt.Println("\n### Comparison of file contents ###")

	for {

		byte1, err := reader1.ReadByte()
		if err != nil && err != io.EOF {
			log.Println("Reading file1 error", err)
			return false, err
		}

		byte2, err := reader2.ReadByte()
		if err != nil && err != io.EOF {
			log.Println("Reading file2 error", err)
			return false, err
		}

		if err == io.EOF && err == io.EOF {
			break
		}

		if byte1 != byte2 {
			fmt.Println("\nFiles differ in content")
			return false, nil
		}
	}

	fmt.Println("\nFiles are identical in content")
	return true, nil
}

func CompareHash(filePath1, filePath2 string) (bool, error) {

	file1, err := os.Open(filePath1)
	if err != nil {
		log.Println("File1 can't be open", err)
		return false, err
	}
	defer file1.Close()

	file2, err := os.Open(filePath2)
	if err != nil {
		log.Println("File2 can't be open", err)
		return false, err
	}
	defer file2.Close()

	hash1 := sha256.New()
	_, err1 := io.Copy(hash1, file1)
	if err1 != nil {
		log.Println("Copy file1 error", err1)
		return false, err1
	}

	hash2 := sha256.New()
	_, err2 := io.Copy(hash2, file2)
	if err2 != nil {
		log.Println("Copy file2 error", err2)
		return false, err2
	}

	fileHash1 := hash1.Sum(nil)
	fileHash2 := hash2.Sum(nil)

	fmt.Println("\n### Comparison of file SHA-256 ###")
	fmt.Printf("\nHash of 1st file: %x\n", fileHash1)
	fmt.Printf("Hash of 2nd file: %x\n", fileHash2)

	for i := range fileHash1 {
		if fileHash1[i] != fileHash2[i] {
			fmt.Println("\nFiles differ in SHA-256")
			return false, nil
		}
	}
	fmt.Println("\nFiles are identical in SHA-256")
	return true, nil
}

func CompareAll(filePath1, filePath2 string) {
	firstCompare, _ := CompareSize(filePath1, filePath2)
	if firstCompare == false {
		fmt.Println("\nFiles differ in content")
		fmt.Println("\nFiles differ in SHA-256")
	} else {
		secondCompare, _ := CompareContent(filePath1, filePath2)
		if secondCompare == false {
			fmt.Println("\nFiles differ in SHA-256")
		} else {
			CompareHash(filePath1, filePath2)
		}
	}
}