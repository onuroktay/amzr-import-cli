package main

import (
	"flag"
	"fmt"
	"github.com/onuroktay/Amazon-Reader/AmzR-pkg-import"
)

func main() {
	fileName := flag.String("file", "amazon.json", "name of the json file to read")
	flag.Parse()

	fmt.Println("Import data from ", *fileName)
	
	err := OnurTPIjsonImporter.ImportJSON(*fileName)
	if err != nil {
		fmt.Println(err.Error())
	}
}
