package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

func main() {

	// var filePath string
	for _,  file := range os.Args[1:] {
		fileInfos, err := ioutil.ReadDir(file)
		fmt.Println(err)
		for _, info := range fileInfos {

			name := info.Name()

			r1 := regexp.MustCompile(".js")
			r2 := regexp.MustCompile(".jsx")
			match1 := r1.MatchString(name)
			match2 := r2.MatchString(name)
			fmt.Println("info",file, name, match1, match2)
			if match1 {
				out := r1.ReplaceAllString(name, ".ts")
				os.Rename(file+"/"+name, file+"/"+out)
			}
			if match2 {
				out := r1.ReplaceAllString(name, ".tsx")
				os.Rename(file+"/"+name, file+"/"+out)
			}
		}

	}


}