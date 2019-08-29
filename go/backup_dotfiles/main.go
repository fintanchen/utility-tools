package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

func main() {

	// Get home dir
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	files, err := ioutil.ReadDir(home)
	if err != nil {
		log.Fatal(err)
	}

	rangeAllFiles(files, home)

}

func rangeAllFiles(files []os.FileInfo, home string) {

	// Range all file in the home directory.
	for _, file := range files {
		reg, err := regexp.Compile("^\\.")
		if err != nil {
			log.Fatal(err)
		}

		if reg.Match([]byte(file.Name())) {
			fmt.Println(file.Name())
			handleDotFile(file, home)
		}
	}

}

func handleDotFile(file os.FileInfo, home string) {
	if !file.IsDir() {
		fmt.Println("Link")
		err := os.Link(home+"/"+file.Name(), "./"+file.Name())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("./" + file.Name())

	}

	if file.IsDir() {

		// err := os.Mkdir("./"+file.Name(), os.ModePerm)
		// if err != nil {
		// 	log.Fatal(err)
		// }

		// err = os.Chdir("./" + file.Name())
		// if err != nil {
		// 	log.Fatal(err)
		// }

		// files, err := ioutil.ReadDir(home + "/" + file.Name())
		// rangeAllFiles(files, true)
	}

}
