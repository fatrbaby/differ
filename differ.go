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

type Differ interface {
	Scan() error
	FileMd5() (string, error)
	Count() int
	Sames() map[string][]string
}

type scanner struct {
	path  string
	Files []string
	sames map[string][]string
	md5   hash.Hash
}

func New(path string) Differ {
	s := &scanner{
		path:  path,
		sames: make(map[string][]string),
		md5:   md5.New(),
	}

	return s
}

func (s *scanner) Scan() error {
	err := filepath.Walk(s.Path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}

		if f.IsDir() {
			return nil
		}

		s.Files = append(s.Files, path)

		return nil
	})

	return err
}

func (s *scanner) FileMd5(file string) (result string, err error) {
	f, err := os.Open(file)

	if err != nil {
		return result, err
	}

	defer f.Close()

	reader := bufio.NewReader(f)
	defer s.md5.Reset()

	if _, err := io.Copy(s.md5, reader); err != nil {
		return result, err
	}

	result = fmt.Sprintf("%x", s.md5.Sum(nil))

	return result, nil
}

func (s *scanner) Count() int {
	return len(s.Files)
}

func (s *scanner) Sames() map[string][]string {
	if len(s.sames) > 0 {
		return s.sames
	}

	counts := make(map[string]int, s.Count())
	hashed := make(map[string][]string)

	for _, file := range s.Files {
		cipher, err := s.FileMd5(file)

		if err == nil {
			counts[cipher]++
			hashed[cipher] = append(hashed[cipher], file)
		} else {
			panic(err)
		}
	}

	for m, c := range counts {
		if c > 1 {
			s.sames[m] = hashed[m]
		}
	}

	return s.sames
}
