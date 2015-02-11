package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	fileName := defaultConfigName

	if len(os.Args) == 2 {
		fileName = os.Args[1]
	}

	p, err := LoadParamsFromFile(fileName)
	if err != nil {

		if err := CreateDefaultConfigFile(); err != nil {
			fmt.Println(err)
			return
		}

		format := "Create a configuration file (\"%s\"). Configure and run the program again\n"
		fmt.Printf(format, defaultConfigName)
		return
	}

	beginT := time.Now()

	if err = CreateFiles(p); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("All successfully completed!")
	fmt.Println("Work time:", time.Since(beginT))
}
