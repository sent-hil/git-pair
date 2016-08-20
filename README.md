# git-pair

This script is a Go port of git-pair binary found in https://github.com/pivotal/git_scripts.

# Installation

* If you're on OSX, move git-pair binary in this repo to somewhere in your $PATH, most probably `/usr/local/bin/`.
* If you're on any other os, build this script with `go build -o git-pair` (you'll need a working Go installation) and move it to somewhere in the $PATH.

# Usage

    $ git pair ij jb
    # => Indiana Jones and James Bond <pair+indiana@gmail.com+james@gmail.com@>

    $ git pair
    # => Indiana Jones indiana@gmail.com

In the root of the folder from which you run this command, there must be a .pairs file.  See .pairs file in this repo for an example.
