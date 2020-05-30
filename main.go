/*
Author: M Krishna
*/
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	args := []string{"."}
	if len(os.Args) > 1 {
		args = os.Args[1:]
	}

	for _, arg := range args {
		err := otherVersion(arg)
		if err != nil {
			log.Printf("Not able to list %s, %v\n", arg, err)
		}
	}
}

// Testing Purpose.
func listCurrentDirectory(root string) error{
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.Name()[0] == '.' {
			return filepath.SkipDir
		}
		rel ,err := filepath.Rel(root, path)

		if err != nil {
			return err
		}

		depth := len(strings.Split(rel, string(filepath.Separator)))

		fmt.Printf("%s%s\n", strings.Repeat("  ", depth), info.Name())
		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

func otherVersion(root string) error {
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
	fmt.Fprint(os.Stdout, "Total Directory: ", totalDirectory, ", Total Files: ", totalFiles)

	return nil
}