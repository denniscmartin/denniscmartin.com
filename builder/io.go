package main

import (
	"fmt"
	"io/fs"
	"os"
	"sort"
	"strings"
)

const IN_INDEX_FILE = "../web/templates/index.html"
const OUT_INDEX_FILE = "../web/out/index.html"
const IN_ARTICLES_DIR = "../web/articles"
const OUT_ARTICLES_DIR = "../web/out/articles"

func readInArticlesDir() []fs.DirEntry {
	files, err := os.ReadDir(IN_ARTICLES_DIR)

	if err != nil {
		panic(err)
	}

	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})

	return files
}

func readInArticle(filename string) string {
	content, err := os.ReadFile(fmt.Sprintf("%s/%s", IN_ARTICLES_DIR, filename))

	if err != nil {
		panic(err)
	}

	return string(content)
}

func readInIndexFile() string {
	content, err := os.ReadFile(IN_INDEX_FILE)

	if err != nil {
		panic(err)
	}

	return string(content)
}

func writeOutIndexFile(content string) {
	err := os.WriteFile(OUT_INDEX_FILE, []byte(content), 0755)

	if err != nil {
		panic(err)
	}
}

func createOutArticle(filename string) (*os.File, string) {
	outFilename := strings.ReplaceAll(filename, ".md", ".html")
	file, err := os.Create(fmt.Sprintf("%s/%s", OUT_ARTICLES_DIR, outFilename))

	if err != nil {
		panic(err)
	}

	return file, fmt.Sprintf("articles/%s", outFilename)
}
