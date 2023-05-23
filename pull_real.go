package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"math/rand"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func listOfFileNames(files []fs.FileInfo) map[string][]string {
	var mp = make(map[string][]string)
	// var ans = make([]string, 0)
	min := 6
	max := 7
	for _, file := range files {
		rand.Seed(time.Now().UnixNano())
		div := rand.Intn(max-min+1) + min

		f := strings.Split(file.Name(), ".")
		// fmt.Println(f[0])
		f1 := strings.Split(f[1], "_")
		v, _ := strconv.Atoi(f1[len(f1)-1])

		if v%div == 0 && v != 0 && len(mp[f[0]]) == 0 && len(mp) <= 400 {
			mp[f[0]] = append(mp[f[0]], file.Name())
			// ans = append(ans, file.Name())
		}
	}
	return mp
}

func copyToPull(mp map[string][]string) {
	for _, files := range mp {
		for _, v := range files {
			in, err := ioutil.ReadFile(filepath.Join("trainB", v))
			if err != nil {
				fmt.Println(err)
				return
			}
			err = ioutil.WriteFile(filepath.Join(".", "real-pull", v), in, 0644)
			if err != nil {
				fmt.Println("Error creating", filepath.Join(".", "real-pull", v))
				fmt.Println(err)
				return
			}
		}
	}

}

func main() {
	files, err := ioutil.ReadDir("trainB")
	if err != nil {
		log.Fatal(err)
	}
	listOfFileNames(files)
	fmt.Println(len(listOfFileNames(files)))

	copyToPull(listOfFileNames(files))
}
