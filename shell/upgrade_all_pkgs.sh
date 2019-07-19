#!/bin/bash

function newface() {
	sleep 1
	clear
}

# for debian-based operate system.
if command -v apt-get >/dev/null 2>&1; then
	apt update && apt dist-upgrade -y
else
	echo "Don't have apt"
fi

newface

# for redhat-based operate system.
if command -v yum >/dev/null 2>&1; then
	yum update -y
else
	echo "Don't have yum"
fi

newface

# for macOS
if command -v brew >/dev/null 2>&1; then
	# update source of brew core and cask
	brew update

	# update brew core
	brew upgrade

	# update brew cask
	brew cask upgrade $(brew cask list | grep -v battle | grep -v steam | grep -v gramm)

	# clean brew install and unused dirs
	brew cleanup
else
	echo "Don't have brew"
fi

newface

# common.
if command -v yarn >/dev/null 2>&1; then
	# yarn
	yarn global upgrade

else
	echo "Don't have yarn"
fi

newface

if command -v rustup >/dev/null 2>&1; then
	# rust
	rustup update
else
	echo "Don't have rustup"
fi

newface

if command -v go >/dev/null 2>&1; then
	# go
	go run $HOME/unsync/github.com/keyvchan/utility-tools/go/update_gopath_repo/upgrade_all_go_pkgs.go
else
	echo "Don't have go"
fi
