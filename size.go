package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var (
	dot       = flag.Bool("dot", false, "Include dot files")
	startPath = flag.String("path", ".", "Set the path to check")
	filePat   = flag.String("files", "*", "Sets a file pattern to use")
)

func main() {
	flag.Parse()

	var files, dirs int64
	var size int64
	t := time.Now()

	err := filepath.Walk(*startPath, func(path string, info os.FileInfo, err error) error {
		if err == nil {
			//log.Print(info.Name(), " ", path)
			if !*dot && path != *startPath && strings.HasPrefix(info.Name(), ".") {
				if info.IsDir() {
					return filepath.SkipDir
				} else {
					return nil
				}
			}
			if !info.IsDir() {
				m, err := filepath.Match(*filePat, info.Name())
				if err != nil {
					log.Print(err)
				}
				if m {
					size += info.Size()
					files++
				}
			} else {
				dirs++
			}
		} else {
			log.Print(err)
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%d folders, %d files, %d bytes, %dK, %dM, %dG\n", dirs, files, size, size/1024, size/1024/1024, size/1024/1024/1024)
	fmt.Printf("Scanned in %s\n", time.Since(t))
}
