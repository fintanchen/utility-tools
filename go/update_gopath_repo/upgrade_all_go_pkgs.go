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
	golang = "golang.org/"
)

func main() {

	GOPATH := os.Getenv("GOPATH")

	// Init WaitGroup.
	wg := sync.WaitGroup{}
	rw := sync.RWMutex{}

	// pull projects in $GOPATH/src/github.com/
	aimDirs := walk(GOPATH, github)

	for _, aim := range aimDirs {
		go pull(&wg, &rw, aim)
	}

	// pull projects in $GOPATH/src/golang.org/
	aimDirs = walk(GOPATH, golang)

	for _, aim := range aimDirs {
		go pull(&wg, &rw, aim)
	}
	wg.Wait()
}

// children get sub files in path.
func children(path string) ([]string, error) {
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

// pull execute git pull in terminal for path.
// Using rw_lock to show reminder and output in same screen.
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
// returning dirs that need to pull
func walk(path string, aim string) []string {

	var pullDirs []string
	parent := path + "/src/" + aim
	authors, err := children(parent)
	if err != nil {
		log.Println(err)
	}

	for _, author := range authors {
		author = parent + author + "/"
		projects, err := children(author)
		if err != nil {
			log.Println(err)
		}

		for _, project := range projects {
			pullDirs = append(pullDirs, author+project+"/")
		}
	}

	return pullDirs
}
