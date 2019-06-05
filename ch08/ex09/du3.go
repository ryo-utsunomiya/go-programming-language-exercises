package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var verbose = flag.Bool("v", false, "show verbose progress messages")

type DirInfo struct {
	root   string
	nfiles int64
	nbytes int64
}

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	result := make(map[string]DirInfo)
	for _, root := range roots {
		result[root] = DirInfo{root: root}
	}

	fileSizes := make(chan DirInfo)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, &n, fileSizes, root)
	}
	go func() {
		n.Wait()
		close(fileSizes)
	}()

	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}

loop:
	for {
		select {
		case f, ok := <-fileSizes:
			if !ok {
				break loop
			}
			c := result[f.root]
			result[f.root] = DirInfo{
				root:   f.root,
				nfiles: c.nfiles + 1,
				nbytes: c.nbytes + f.nbytes,
			}
		case <-tick:
			printDiskUsage(result)
		}
	}
	printDiskUsage(result)
}

func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- DirInfo, root string) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, n, fileSizes, root)
		} else {
			fileSizes <- DirInfo{
				root:   root,
				nbytes: entry.Size(),
			}
		}
	}
}

var sema = make(chan struct{}, 20)

func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}
	defer func() { <-sema }()
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}

func printDiskUsage(result map[string]DirInfo) {
	for k, v := range result {
		fmt.Printf("%s: %d files %.1f GB\n", k, v.nfiles, float64(v.nbytes)/1e9)
	}
}
