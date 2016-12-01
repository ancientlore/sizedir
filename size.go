package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

var (
	dot       = flag.Bool("dot", false, "Include dot files")
	startPath = flag.String("path", ".", "Set the path to check")
	filePat   = flag.String("files", "*", "Sets a file pattern to use")
	ext       = flag.Bool("ext", false, "Aggregate by extension")
)

type stats struct {
	files int64
	size  int64
}

func main() {
	flag.Parse()

	var files, dirs int64
	var size int64
	t := time.Now()
	extensions := make(map[string]stats)

	err := filepath.Walk(*startPath, func(path string, info os.FileInfo, err error) error {
		if err == nil {
			//log.Print(info.Name(), " ", path)
			if !*dot && path != *startPath && strings.HasPrefix(info.Name(), ".") {
				if info.IsDir() {
					return filepath.SkipDir
				}
				return nil
			}
			if !info.IsDir() {
				m, err := filepath.Match(*filePat, info.Name())
				if err != nil {
					log.Print(err)
				}
				if m {
					size += info.Size()
					files++
					if *ext {
						extn := filepath.Ext(info.Name())
						s := extensions[extn]
						s.files++
						s.size += info.Size()
						extensions[extn] = s
					}
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

	fmt.Printf("%d folders, %d files, %d bytes, %dKB, %dMB, %dGB\n", dirs, files, size, size/1024, size/1024/1024, size/1024/1024/1024)
	fmt.Printf("Scanned in %s\n", time.Since(t))
	if *ext {
		fmt.Println()
		var keys []string
		max := 3
		maxc := 5
		maxs := 5
		maxk := 2
		maxm := 2
		maxg := 2
		maxa := 6
		for k := range extensions {
			keys = append(keys, k)
			if len(k) > max {
				max = len(k)
			}
			digits := len(strconv.FormatInt(extensions[k].files, 10))
			if digits > maxc {
				maxc = digits
			}
			digits = len(strconv.FormatInt(extensions[k].size, 10))
			if digits > maxs {
				maxs = digits
			}
			digits = len(strconv.FormatInt(extensions[k].size/1024, 10))
			if digits > maxk {
				maxk = digits
			}
			digits = len(strconv.FormatInt(extensions[k].size/1024/1024, 10))
			if digits > maxm {
				maxm = digits
			}
			digits = len(strconv.FormatInt(extensions[k].size/1024/1024/1024, 10))
			if digits > maxg {
				maxg = digits
			}
			digits = len(strconv.FormatInt(extensions[k].size/extensions[k].files, 10))
			if digits > maxa {
				maxa = digits
			}
		}
		sort.Strings(keys)
		fmt.Printf("%-*.*s %*s %*s %*s %*s %*s %*.*s\n", max, max, "Extension", maxc, "Files", maxs, "Bytes", maxk, "KB", maxm, "MB", maxg, "GB", maxa, maxa, "Avg KB")
		fmt.Printf("%s %s %s %s %s %s %s\n", strings.Repeat("-", max), strings.Repeat("-", maxc), strings.Repeat("-", maxs), strings.Repeat("-", maxk), strings.Repeat("-", maxm), strings.Repeat("-", maxg), strings.Repeat("-", maxa))

		for _, k := range keys {
			s := extensions[k]
			fmt.Printf("%-*s %*d %*d %*d %*d %*d %*d\n", max, k, maxc, s.files, maxs, s.size, maxk, s.size/1024, maxm, s.size/1024/1024, maxg, s.size/1024/1024/1024, maxa, s.size/s.files)
		}
	}
}
