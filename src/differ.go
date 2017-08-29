package main

import (
	"fmt"
	"os"
	"path/filepath"
	"crypto/md5"
	"io"
	"bufio"
	"flag"
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

	for _, file := range Files {
		md5, err := FileMd5(file)

		if err == nil {
			fmt.Println(md5)
		} else {
			fmt.Printf("%v\n", err)
		}
	}
}

func Scan(path string) error {
	err := filepath.Walk(path, func (path string, f os.FileInfo, err error) error {
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
