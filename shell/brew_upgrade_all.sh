#!/bin/bash

# update source of brew core and cask
brew update

# update brew core
brew upgrade

# update brew cask
brew cask upgrade `brew cask list`

# clean brew install and unused dirs
brew cleanup && brew prune

# yarn
yarn global upgrade

# rust
rustup update

# go
go run $HOME/unSync/utility-tools/go/upgrade_all_go_pkgs.go

