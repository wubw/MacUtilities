package main

import (
    "fmt"
	"os"
	"path"
	"path/filepath"
	"log"
)

func main() {
	var files []string

    root := "/Users/wubinwei/Downloads/Photo Exported"
    err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
        files = append(files, path)
        return nil
    })
    if err != nil {
        panic(err)
	}
    for _, f := range files {
		fmt.Println(f)
		
		fi, err := os.Stat(f)
		if err != nil {
			fmt.Println(err)
			return
		}
		switch mode := fi.Mode(); {
		case mode.IsDir():
		case mode.IsRegular():
			dir, file := filepath.Split(f)
			if file == ".DS_Store" {
				break
			}
			fmt.Println("rename to => ")
			newname := path.Join(root, filepath.Base(dir)+" "+file)
			fmt.Println(newname)
			fmt.Println()
			err := os.Rename(f, newname)
			if err != nil {
				log.Fatal(err)
			}
		}
    }
}
