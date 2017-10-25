package differ

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type Differ struct {
	Files []string
}

func (d *Differ) Scan(path string) error {
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}

		if f.IsDir() {
			return nil
		}

		d.Files = append(d.Files, path)

		return nil
	})

	return err
}

func (d *Differ) FileMd5(file string) (result string, err error) {
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

func (d *Differ) Count() int  {
	return len(d.Files)
}

func (d *Differ) Sames() map[string][]string {
	counts := make(map[string]int, d.Count())
	hashed := make(map[string][]string)
	sames := make(map[string][]string)

	for _, file := range d.Files {
		cipher, err := d.FileMd5(file)

		if err == nil {
			counts[cipher]++
			hashed[cipher] = append(hashed[cipher], file)
		} else {
			panic(err)
		}
	}

	for m, c := range counts {
		if c > 1 {
			sames[m] = hashed[m]
		}
	}

	return sames
}
