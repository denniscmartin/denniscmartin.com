package main

import (
	"fmt"
	"io/fs"
	"os"
	"sort"
	"strings"
)

const indexFilename = "../web/src/index.html"
const articlesPath = "../web/articles"
const outputPath = "../web/src/articles"

func readInputArticlesDir() []fs.DirEntry {
	inputArticles, err := os.ReadDir(articlesPath)

	if err != nil {
		panic(err)
	}

	sort.Slice(inputArticles, func(i, j int) bool {
		return inputArticles[i].Name() < inputArticles[j].Name()
	})

	return inputArticles
}

func openInputArticle(filename string) *os.File {
	inputArticle, err := os.Open(fmt.Sprintf("%s/%s", articlesPath, filename))

	if err != nil {
		panic(err)
	}

	return inputArticle
}

func openIndex() *os.File {
	file, err := os.Open(indexFilename)

	if err != nil {
		panic(err)
	}

	return file
}

func createOutputArticle(filename string) (*os.File, string) {
	outputFilename := strings.ReplaceAll(filename, ".md", ".html")
	outputArticle, err := os.Create(fmt.Sprintf("%s/%s", outputPath, outputFilename))

	if err != nil {
		panic(err)
	}

	return outputArticle, fmt.Sprintf("articles/%s", outputFilename)
}
