package main

import (
	"bufio"
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

var Files []string

func main() {
	directory := flag.String("dir", "/null", "The directory you want to scan")
	flag.Parse()

	if _, err := os.Stat(*directory); err != nil {
		panic(err)
	}

	err := Scan(*directory)

	if err != nil {
		panic(err)
	}

	counts := make(map[string]int)

	for _, file := range Files {
		md5, err := FileMd5(file)

		if err == nil {
			counts[md5]++
		} else {
			fmt.Printf("%v\n", err)
		}
	}
	
	fmt.Printf("%d, %d\n", len(Files), len(counts))
}

func Scan(path string) error {
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}

		if f.IsDir() {
			return nil
		}

		Files = append(Files, path)

		return nil
	})

	return err
}

func FileMd5(file string) (result string, err error) {
	f, err := os.Open(file)

	if err != nil {
		return result, err
	}

	defer f.Close()

	reader := bufio.NewReader(f)
	hasher := md5.New()

	if _, err := io.Copy(hasher, reader); err != nil {
		return result, err
	}

	result = fmt.Sprintf("%x", hasher.Sum(nil))

	return result, nil
}
