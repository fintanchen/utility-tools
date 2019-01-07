package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"sync"
)

var (
	github = "github.com/"
)

func main() {

	GOPATH := os.Getenv("GOPATH")

	// Init WaitGroup.
	wg := sync.WaitGroup{}
	rw := sync.RWMutex{}

	aimDirs := walk(GOPATH)

	for _, aim := range aimDirs {
		go pull(&wg, &rw, aim)
	}

	wg.Wait()
}

// subFiles get sub files in path.
func subFiles(path string) ([]string, error) {
	p, err := os.Open(path)
	if err != nil {
		log.Println(err)
	}
	subFs, err := p.Readdirnames(0)
	if err != nil {
		log.Println(err)
	}

	return subFs, err
}

// pull execute git pull
func pull(wg *sync.WaitGroup, rw *sync.RWMutex, path string) {

	wg.Add(1)
	defer wg.Done()

	rw.Lock()
	fmt.Println("Starting pull", path)

	err := os.Chdir(path)
	if err != nil {
		log.Println(err)
	}

	cmd := exec.Command("git", "pull")
	o, err := cmd.CombinedOutput()
	if err != nil {
		log.Println("Couldn't run ", err)
	}
	fmt.Println(string(o))

	rw.Unlock()
}

// walk through the path.
func walk(path string) []string {

	var pullDirs []string
	github := path + "/src/github.com/"
	authors, err := subFiles(github)
	if err != nil {
		log.Println(err)
	}

	for _, dir := range authors {
		dir = github + dir
		reals, err := subFiles(dir + "/")
		if err != nil {
			log.Println(err)
		}

		for _, real := range reals {
			pullDirs = append(pullDirs, dir+"/"+real+"/")
		}
	}

	return pullDirs
}
