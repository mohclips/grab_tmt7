package main

import (
	"fmt"
	"log"

	"github.com/whitone/clickonce"
)

var manifest = "https://aka.ms/threatmodelingtool"
var download_dir = "downloaded_files"

func main() {
	fmt.Printf("Please wait a few minutes while %s is downloaded\n", manifest)

	co := &clickonce.ClickOnce{}

	co.SetOutputDir(download_dir)

	co.Init(manifest)

	err := co.GetAll()
	if err != nil {
		log.Fatal(err)
	}

	// for dPath, dContent := range co.DeployedFiles() {
	// 	bcont := []byte(dContent.Content)
	// 	fmt.Printf("%s: %s\n", dPath, http.DetectContentType(bcont))
	// }

	fmt.Printf("Downloaded\nCheck %s for your files\n", download_dir)

}
