# ls
For listing directories and files. Windows (Powershell) and linux have this command, I thought of recreating it.

## Download
`go get github.com/M-krishna/ls`

## Install the Package
`go install`

## Setup the environment variable for Windows
So that you can access the `ls` command outside of your `$GOPATH`

## Usage
* Type `ls` to get the directories and files from current path
* Type `ls ../folder ../../folder` for multiple folders
