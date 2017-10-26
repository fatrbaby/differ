package differ

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"hash"
	"io"
	"os"
	"path/filepath"
)

type Differ struct {
	Path      string
	Files     []string
	sames     map[string][]string
	md5 hash.Hash
}

func New(directory string) *Differ {
	d := &Differ{
		Path:      directory,
		sames:     make(map[string][]string),
		md5: md5.New(),
	}

	return d
}

func (d *Differ) Scan() error {
	err := filepath.Walk(d.Path, func(path string, f os.FileInfo, err error) error {
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
	d.md5.Reset()

	if _, err := io.Copy(d.md5, reader); err != nil {
		return result, err
	}

	result = fmt.Sprintf("%x", d.md5.Sum(nil))

	return result, nil
}

func (d *Differ) Count() int {
	return len(d.Files)
}

func (d *Differ) Sames() map[string][]string {
	if len(d.sames) > 0 {
		return d.sames
	}

	counts := make(map[string]int, d.Count())
	hashed := make(map[string][]string)

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
			d.sames[m] = hashed[m]
		}
	}

	return d.sames
}
