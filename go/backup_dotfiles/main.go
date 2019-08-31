package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {

	dst := os.Args[1]

	// Get home dir
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	files, err := ioutil.ReadDir(home)
	if err != nil {
		log.Fatal(err)
	}

	dotFiles := rangeAllFiles(files)

	for _, dotFile := range dotFiles {
		if dotFile.IsDir() {
			err = filepath.Walk(home+"/"+dotFile.Name(), func(path string, info os.FileInfo, err error) error {
				if err != nil {
					fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
					return err
				}

				trimEd := strings.TrimPrefix(path, home)
				if info.IsDir() {

					err := os.Mkdir(dst+trimEd, os.ModePerm)
					if err != nil {
						log.Fatal(err)
					}

					return nil
				}

				fmt.Println(dst + trimEd)
				err = os.Link(path, dst+trimEd)
				if err != nil {
					log.Fatal(err)
				}

				fmt.Printf("Linked file: %q\n", dst+trimEd)

				return nil
			})

		} else {

			err := os.Link(home+"/"+dotFile.Name(), dst+dotFile.Name())
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Linked file: %q\n", dst+dotFile.Name())
		}

	}

}

func rangeAllFiles(files []os.FileInfo) (dotFiles []os.FileInfo) {

	var dotfiles = []os.FileInfo{}

	// Range all file in the home directory.
	for _, file := range files {
		reg, err := regexp.Compile("^\\.")
		if err != nil {
			log.Fatal(err)
		}

		if reg.Match([]byte(file.Name())) {
			fmt.Println(file.Name())
			dotfiles = append(dotfiles, file)
		}
	}

	// Get all dot files
	return dotfiles

}
