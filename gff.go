package main

import "fmt"
import "io/ioutil"
import "regexp"
import "flag"

var re = regexp.MustCompile(".w.$")

func readDirectory(dir string, depth int) []string {
	// fmt.Printf("%s with depth %d\n", dir, depth)
	if depth < 0 {
		return []string{}
	}

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		//fmt.Printf("Failed to open directory %s: %s\n", dir, err)
		return []string{}
	}

	out := []string{}
	for _, file := range files {
		if re.MatchString(fmt.Sprintf("%s", file.Mode())) {
			out = append(out, fmt.Sprintf("%s %s", dir+"/"+file.Name(), file.Mode()))
			fmt.Println(fmt.Sprintf("%s %s", dir+"/"+file.Name(), file.Mode()))
		}
	}

	for _, file := range files {
		if file.IsDir() {
			out = append(out, readDirectory(dir+"/"+file.Name(), depth-1)...)
		}
	}

	return out
}

func main() {
	fmt.Println("Global File Finder")

	dirPtr := flag.String("dir", ".", "directory to start searching from")
	depthPtr := flag.Int("depth", 3, "depth of directories to search")
	flag.Parse()

	readDirectory(*dirPtr, *depthPtr)
}
