package main

import (
	"fmt"
	"os"
	"path/filepath"
)

var Files []string

func main() {
	/*
	wc, _ := os.Getwd()
	var directory = flag.String("dir", wc, "which dir you wanna scan")
	flag.Parse()
	*/

	err := Scan("/usr/local/var/www/app")

	if err == nil {
		fmt.Println("files:")
		fmt.Sprintf("%v\n", Files)
	} else {
		fmt.Println("errors:")
		fmt.Fprintf(os.Stderr, "%v", err)
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
