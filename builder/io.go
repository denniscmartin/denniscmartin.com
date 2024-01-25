package main

import (
	"fmt"
	"io/fs"
	"os"
	"sort"
	"strings"
)

const inputIndexFilename = "../web/templates/index.html"
const outputIndexFilename = "../web/out/index.html"
const inputArticlesDir = "../web/articles"
const outputArticlesDir = "../web/out/articles"

func readInputArticlesDir() []fs.DirEntry {
	files, err := os.ReadDir(inputArticlesDir)

	if err != nil {
		panic(err)
	}

	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})

	return files
}

func openInputArticle(filename string) *os.File {
	file, err := os.Open(fmt.Sprintf("%s/%s", inputArticlesDir, filename))

	if err != nil {
		panic(err)
	}

	return file
}

func readInputIndex() string {
	content, err := os.ReadFile(inputIndexFilename)

	if err != nil {
		panic(err)
	}

	return string(content)
}

func writeOutputIndex(content string) {
	err := os.WriteFile(outputIndexFilename, []byte(content), 0755)

	if err != nil {
		panic(err)
	}
}

func createOutputArticle(filename string) (*os.File, string) {
	outFilename := strings.ReplaceAll(filename, ".md", ".html")
	file, err := os.Create(fmt.Sprintf("%s/%s", outputArticlesDir, outFilename))

	if err != nil {
		panic(err)
	}

	return file, fmt.Sprintf("articles/%s", outFilename)
}
