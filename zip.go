package main

import (
	"flag"
	"fmt"
	"os"
	//"io"
	"io/ioutil"
	//"bytes"
	"archive/zip"
	"path/filepath"
	"strings"
)

func addFilesToZip(root string) error {
	fw, _ := os.Create("test.zip")
	defer fw.Close()
	// Create a new zip archive.
	writer := zip.NewWriter(fw)
	defer writer.Close()

	walk := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		if err != nil {
			fmt.Println("error open file")
			return nil
		}

		strRoots := strings.Split(root, "/")
		strPaths := strings.Split(path, "\\")

		dst := ""
		for i := len(strRoots); i < len(strPaths); i++ {
			if i != len(strPaths)-1 {
				dst = dst + strPaths[i] + "/"
			} else {
				dst = dst + strPaths[i]
			}
		}

		zip, err := writer.Create(dst)
		if err != nil {
			fmt.Println("error create")
			return nil
		}

		data, err := ioutil.ReadFile(path)
		_, err = zip.Write(data)

		if err != nil {
			fmt.Println(err)
		}

		return nil
	}

	return filepath.Walk(root, walk)
}

func main() {
	flag.Parse()
	var src string

	if flag.NArg() > 0 {
		src = flag.Arg(0)
	} else {
		fmt.Println("Provide the src directory!")
		return
	}

	addFilesToZip(src)
}
