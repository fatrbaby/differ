package main

import (
	"path/filepath"
	"os"
	"runtime"
	"fmt"
	"crypto/md5"
	"bufio"
	"io"
)

type filepack struct {
	Code string
	Name string
}

func main() {
	Sames("/usr/local/var/www")
}

func extractMd5(files [][]string)  <-chan filepack {
	out := make(chan filepack)
	hasher := md5.New()

	for _, file := range files {
		go func() {
			hasher.Reset()
			for _, f := range file {
				fi, err := os.Open(f)

				if err != nil {
					panic(err)
				}
				reader := bufio.NewReader(fi)

				if _, err := io.Copy(hasher, reader); err != nil {
					panic(err)
				}
				code := fmt.Sprintf("%x", hasher.Sum(nil))
				out <- filepack{Code: code, Name:f}
			}
		}()
	}

	close(out)

	return out
}

func chunks(files []string, size int) [][]string {
	t := 0
	chunk := make([][]string, size)

	for _, f := range files {
		odd := t %  size
		chunk[odd] = append(chunk[odd], f)
		t++
	}

	return chunk
}

func scan(dir string) ([]string, error) {
	var files []string

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info == nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		files = append(files, info.Name())

		return nil
	})

	return files, err
}

func Sames(dir string) {
	files, err := scan(dir)

	if err != nil {
		panic(err)
	}

	chunked := chunks(files, runtime.NumCPU())

	codes := extractMd5(chunked)
	counts := make(map[string]int)
	hashed := make(map[string][]string)
	sames := make(map[string][]string)
	t := 0
	for code := range codes {
		counts[code.Code]++
		hashed[code.Code] = append(hashed[code.Code], code.Name)
		t++
	}

	for m, c := range counts {
		if c > 1 {
			if c > 1 {
				sames[m] = hashed[m]
			}
		}
	}

	fmt.Printf("%v\n", t)
}
