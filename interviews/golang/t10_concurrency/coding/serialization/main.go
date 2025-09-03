// Example:
// go run main.go .
// d41d8cd98f00b204e9800998ecf8427e  README.md
// a14ed51662ad556474e0295b5db96539  main.go
package main

import (
	"crypto/md5"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"sync"
)

func main() {
	// Calculate the MD5 sum of all files under the specified directory,
	// then print the results sorted by path name.
	m, err := MD5All(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	var paths []string
	for path := range m {
		paths = append(paths, path)
	}
	sort.Strings(paths)
	for _, path := range paths {
		fmt.Printf("%x  %s\n", m[path], path)
	}
}

// MD5All reads all the files in the file tree rooted at root and returns a map
// from file path to the MD5 sum of the file's contents.  If the directory walk
// fails or any read operation fails, MD5All returns an error.
func MD5All(root string) (map[string][md5.Size]byte, error) {
	done := make(chan struct{})
	defer close(done)

	walkResults := walkFiles(done, root)

	results := make(chan result)

	var wg sync.WaitGroup
	wg.Add(digestersNum)
	for i := 0; i < digestersNum; i++ {
		go func() {
			defer wg.Done()
			digester(done, walkResults, results)
		}()
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	m := make(map[string][md5.Size]byte)
	for r := range results {
		if r.err != nil {
			return nil, r.err
		}
		m[r.path] = r.sum
	}

	return m, nil
}

const digestersNum = 5

type result struct {
	path string
	sum  [md5.Size]byte
	err  error
}

type walkResult struct {
	path string
	err  error
}

func walkFiles(done <-chan struct{}, root string) <-chan walkResult {
	results := make(chan walkResult)

	go func() {
		defer func() {
			close(results)
		}()
		err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.Mode().IsRegular() {
				return nil
			}

			select {
			case results <- walkResult{path: path}:
				fmt.Println("walk path", path)
			case <-done:
				return errors.New("walk cancelled")
			}

			return nil
		})
		if err != nil {
			results <- walkResult{err: err}
		}
	}()

	return results
}

func digester(done <-chan struct{}, walkResults <-chan walkResult, out chan<- result) {
	for wr := range walkResults {
		if wr.err != nil {
			return
		}

		fmt.Println("digester path", wr.path)
		var r result
		data, err := os.ReadFile(wr.path)
		if err != nil {
			r = result{
				path: wr.path,
				err:  err,
			}
		} else {
			r = result{
				path: wr.path,
				sum:  md5.Sum(data),
			}
		}

		select {
		case <-done:
			fmt.Println("digester cancelled")
			return
		case out <- r:
			fmt.Println("digester result", r)
		}
	}
}
