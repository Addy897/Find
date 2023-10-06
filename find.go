package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Print("[-]Usage: find.exe [Directory To Search] [Filename]")
		return
	} else {

		dir := os.Args[1]
		file := os.Args[2]
		count:=0
		countFiles(dir, file,&count)
		fmt.Printf("Total Files Found: %d\n",count)
	}

}
func countFiles(name string, file string,count *int) {
	files, err := os.ReadDir(name)
	if err != nil {
		//fmt.Println(err.Error())
		//wg.Done()
		return
	}
	for _, v := range files {
		currentDir := name + v.Name() + "\\"
		if !v.IsDir() {

			if strings.Contains(v.Name(),file) {
				fmt.Printf("Found %s --> %s\n", file, name+v.Name())
				*count++
			}

		} else {
			countFiles(currentDir, file,count)

		}
	}

	return
}
