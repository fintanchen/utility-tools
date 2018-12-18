#!/bin/bash

# update source of brew core and cask
brew update

# update brew core
brew upgrade

# update brew cask
brew cask upgrade `brew cask list`

# clean brew install and unused dirs
brew cleanup && brew prune


