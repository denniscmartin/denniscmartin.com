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

func readFile(filename string) string {
	content, err := os.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	return string(content)
}

func writeFile(content string, filename string) {
	err := os.WriteFile(filename, []byte(content), 0755)

	if err != nil {
		panic(err)
	}
}
