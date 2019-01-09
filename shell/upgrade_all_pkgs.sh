#!/bin/bash

if  command -v apt-get > /dev/null 2>&1;then
    apt update && apt dist-upgrade -y
fi

if  command -v yum > /dev/null 2>&1;then
    yum update -y
fi

if  command -v brew > /dev/null 2>&1;then
    # update source of brew core and cask
    brew update

    # update brew core
    brew upgrade

    # update brew cask
    brew cask upgrade `brew cask list`

    # clean brew install and unused dirs
    brew cleanup && brew prune
fi


if  command -v yarn > /dev/null 2>&1;then
    # yarn
    yarn global upgrade

fi


if  command -v rustup > /dev/null 2>&1;then
    # rust
    rustup update
fi

if  command -v go > /dev/null 2>&1;then
    # go
    go run $HOME/unSync/utility-tools/go/update_gopath_repo/upgrade_all_go_pkgs.go
fi

