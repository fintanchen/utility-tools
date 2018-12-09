package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
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

	ctx, cf := context.WithCancel(context.Background())

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
					go gitPull(ctx, name)

				}
			}
		}

	}

	select {
	case <-ctx.Done():
		cf()
	}
}

func getSubf(path string) ([]string, error) {
	p, err := os.Open(path)
	subFs, err := p.Readdirnames(0)
	return subFs, err
}

func gitPull(ctx context.Context, path string) {

	err := os.Chdir(path)
	if err != nil {
		log.Println(err)
	}

	cmd := exec.Command("git", "pull")
	o, err := cmd.CombinedOutput()
	if err != nil {
		log.Println("Couldn't run ", err)
	}
	fmt.Println("Starting pull", path)
	fmt.Println(string(o))
	<-ctx.Done()
}
