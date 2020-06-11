/*
Author: M Krishna
*/
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	args := []string{"."}
	if len(os.Args) > 1 {
		args = os.Args[1:]
	}

	for _, arg := range args {
		err := listDirectory(arg)
		if err != nil {
			log.Printf("Not able to list %s, %v\n", arg, err)
		}
	}
}

func listDirectory(root string) error {
	fi, err := os.Stat(root)

	if err != nil {
		return err
	}

	if !fi.IsDir() {
		return nil
	}

	fis, err := ioutil.ReadDir(root)

	if err != nil {
		return err
	}

	var totalDirectory int = 0
	var totalFiles int = 0

	for _, info := range fis {
		if info.Name()[0] != '.' && info.Name()[0] != '$'{
			if info.IsDir() {
				totalDirectory += 1
			} else {
				totalFiles += 1
			}
			fmt.Printf(" %v\n", info.Name())
		}
	}
	fmt.Fprint(os.Stdout, "\n")
	fmt.Fprint(os.Stdout, "Total Directory: ", totalDirectory, ", Total Files: ", totalFiles)

	return nil
}