package main

import (
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

	subDirs, err := getSubf(GOPATH + "/src/")
	if err != nil {
		log.Fatal("src", err)
	}

	wg := sync.WaitGroup{}

	for _, dir := range subDirs {
		if dir == "github.com" || dir == "golang.org" {

			var pullPath string
			if dir == "github.com" {
				pullPath = GOPATH + "/src/" + github
			} else {
				pullPath = GOPATH + "/src/" + golangx

			}

			userNames, err := getSubf(pullPath)
			if err != nil {
				log.Fatal("github.com", err)
			}

			for _, name := range userNames {

				if name == "" {
					break
				}
				repoNames, err := getSubf(pullPath + name)
				if err != nil {
					log.Println(repoNames, " ", err)
				}

				for _, repoName := range repoNames {
					name := pullPath + name + "/" + repoName
					go gitPull(&wg, name)

				}
			}
		}

	}

	wg.Wait()
}

func getSubf(path string) ([]string, error) {
	p, err := os.Open(path)
	subFs, err := p.Readdirnames(0)
	return subFs, err
}

func gitPull(wg *sync.WaitGroup, path string) {

	wg.Add(1)
	defer wg.Done()

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
}
