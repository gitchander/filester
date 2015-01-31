package main

import (
	"fmt"
)

func main() {

	fileName := defaultConfigName

	if len(os.Args) == 2 {
		fileName = os.Args[1]
	}

	p, err := loadParamsFromFile(fileName)
	if err != nil {

		if err := createDefaultConfigFile(); err != nil {
			fmt.Println(err)
			return
		}

		format := "Create a configuration file (\"%s\"). Configure and run the file again\n"
		fmt.Printf(format, defaultConfigName)
		return
	}

	if err = createFiles(p); err != nil {
		fmt.Println(err)
		return
	}
}
