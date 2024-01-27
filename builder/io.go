package main

import (
	"io/fs"
	"os"
	"sort"
)

func readDir(path string) []fs.DirEntry {
	files, err := os.ReadDir(path)

	if err != nil {
		panic(err)
	}

	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})

	return files
}

func createDir(path string) {
	err := os.Mkdir(path, 0755)

	if err != nil {
		panic(err)
	}
}

func deleteDir(path string) {
	err := os.RemoveAll(path)

	if err != nil {
		panic(err)
	}
}

func readFile(filename string) []byte {
	content, err := os.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	return content
}

func writeFile(content []byte, filename string) {
	err := os.WriteFile(filename, content, 0755)

	if err != nil {
		panic(err)
	}
}
