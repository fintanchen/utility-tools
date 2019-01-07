package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"sync"
)

var (
	github  = "github.com/"
	golangx = "golang.org/"
)

func main() {
	GOPATH := os.Getenv("GOPATH")

	// Init WaitGroup.
	wg := sync.WaitGroup{}
	rw := sync.RWMutex{}

	githubSubDirs, err := getSubf(GOPATH + "/src/github.com/")
	if err != nil {
		log.Fatal(errors.New("Get sub dir: Can't get sub dir."))
	}

	for _, dir := range githubSubDirs {
		aim, err := getSubf(GOPATH + "/src/github.com/" + dir + "/")
		if err != nil {
			log.Println(errors.New("Range github sub dits: Can't get sub dirs"))
		}
		for _, aimd := range aim {
			aimd = GOPATH + "/src/github.com/" + dir + "/" + aimd + "/"
			gitPull(&wg, aimd, &rw)
		}
	}

	wg.Wait()
}

func getSubf(path string) ([]string, error) {
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

func gitPull(wg *sync.WaitGroup, path string, rw *sync.RWMutex) {

	rw.Lock()
	wg.Add(1)
	defer wg.Done()
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
	//	fmt.Println("Starting pull", path)
	fmt.Println(string(o))
	rw.Unlock()
}
