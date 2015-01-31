package main

import (
	"fmt"
	fst "github.com/gitchander/filester"
	"os"
)

func main() {

	fileName := fst.DefaultConfigName

	if len(os.Args) == 2 {
		fileName = os.Args[1]
	}

	p, err := fst.LoadParamsFromFile(fileName)
	if err != nil {

		if err := fst.CreateDefaultConfigFile(); err != nil {
			fmt.Println(err)
			return
		}

		format := "Create a configuration file (\"%s\"). Configure and run the file again\n"
		fmt.Printf(format, fst.DefaultConfigName)
		return
	}

	if err = fst.CreateFiles(p); err != nil {
		fmt.Println(err)
		return
	}
}
