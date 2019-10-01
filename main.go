package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func dirTree(path string, lvl int) error {
	lvl++
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	tabs := ""
	for i := 0; i < lvl; i++ {
		tabs += "\t"
	}

	for _, file := range files {
		if file.IsDir() {
			fmt.Println(tabs + " - " + file.Name())
			dirTree(path+"/"+file.Name(), lvl)
		} else {
			fmt.Println(tabs + file.Name())
		}
	}

	return nil
}

func main() {

	fmt.Print("Enter path: ")
	path := ""
	fmt.Scanln(&path)

	dirTree(path, 0)
}
