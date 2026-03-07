package main

import (
	"fmt"
	"path/filepath"
	"io/fs"
	"sort"
	"flag"
	"strings"
)

type ExtStat struct {
	Ext string
	Size int64
}

func findFiles(root string, includeHidden bool) (map[string]int64, error) {
	state := make(map[string]int64)
	
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("Failed to walk directory: %w", err)
		}

		if !includeHidden && strings.HasPrefix(d.Name(), ".") {
			if d.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		if d.Type() & fs.ModeSymlink != 0 {
			return nil
		}

		if !d.IsDir() {
			info, err := d.Info()
			if err != nil {
				fmt.Errorf("Failed to get file info: %w", err)
				return nil
			}
			ext := filepath.Ext(path)
			if ext == ""{
				ext = "(no ext)"
			}
			state[ext] += info.Size()
		}
		return nil
	})
	return state, err

}

func main(){
	SearchDir := "/Users/ogawadaichi/Desktop/practice/golang/go-tools/"
	includeHidden := flag.Bool("hidden", false, "Include hidden files in the search")
	flag.Parse()

	extMap, err := findFiles(SearchDir, *includeHidden)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	var stats []ExtStat
	for ext, size := range extMap {
		stats = append(stats, ExtStat{ext, size})
	}

	sort.Slice(stats, func(i, j int) bool {
		return stats[i].Size > stats[j].Size
	})

	fmt.Printf("--- Top 10 file extensions by total size ---\n")
	for i := 0; i < len(stats) && i < 10; i++ {
		fmt.Printf("%2d. [%-10s] %12d bytes\n", i+1, stats[i].Ext, stats[i].Size)
	}
}