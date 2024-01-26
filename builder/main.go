package main

import (
	"bufio"
	"sort"
	"strings"
)

type ArticleItem struct {
	Title       string
	Description string
	Date        string
	Url         string
}

func main() {
	var articles []ArticleItem
	inputArticlesDir := readInputArticlesDir()

	sort.Slice(inputArticlesDir, func(i, j int) bool {
		return inputArticlesDir[i].Name() > inputArticlesDir[j].Name()
	})

	for _, inputArticleEntry := range inputArticlesDir {
		outputArticle, url := createOutputArticle(inputArticleEntry.Name())
		defer outputArticle.Close()

		outputArticle.WriteString(openHtml)
		outputArticle.WriteString(head)
		outputArticle.WriteString(openBody)
		outputArticle.WriteString(navbar)
		outputArticle.WriteString(openMain)

		inputArticle := openInputArticle(inputArticleEntry.Name())
		defer inputArticle.Close()

		articleData := ArticleItem{Url: url}
		inputArticleScanner := bufio.NewScanner(inputArticle)
		readingMetadata := true

		for inputArticleScanner.Scan() {
			lineStr := inputArticleScanner.Text()

			if lineStr == METADATA_END {
				readingMetadata = false
				continue
			}

			// Parse metadata
			if readingMetadata {
				iEqual := strings.Index(lineStr, "=")

				if iEqual == -1 {
					panicInvalidLine(lineStr, inputArticleEntry.Name())
				}

				value := lineStr[iEqual+1:]

				switch tag := lineStr[:iEqual]; tag {
				case METADATA_TITLE:
					articleData.Title = value
				case METADATA_DESCR:
					articleData.Description = value
				case METADATA_DATE:
					articleData.Date = value
				default:
					panicInvalidLine(lineStr, inputArticleEntry.Name())
				}
			} else {
				// Parse markdown body
				lineRunes := []rune(lineStr)

				for i, r := range lineRunes {
					switch r {
					case '#':

					}
				}

			}
		}

		outputArticle.WriteString(closeMain)
		outputArticle.WriteString(closeBody)
		outputArticle.WriteString(closeHtml)

		articles = append(articles, articleData)

		var htmlArticlesList string

		for _, article := range articles {
			htmlArticlesList += createArticleListItem(article)
		}

		indexContent := readInputIndex()
		newIndexContent := strings.ReplaceAll(indexContent, "{{ARTICLES}}", htmlArticlesList)

		if indexContent == newIndexContent {
			panic("articles list not generated")
		}

		writeOutputIndex(newIndexContent)
	}
}
