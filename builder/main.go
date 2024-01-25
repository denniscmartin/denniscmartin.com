package main

import (
	"bufio"
	"strings"
)

type ArticleItem struct {
	Description string
	Date        string
	Url         string
}

func main() {
	var articles []ArticleItem
	inputArticlesDir := readInputArticlesDir()

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
		readingMetadata := false

		for inputArticleScanner.Scan() {
			lineStr := inputArticleScanner.Text()

			if lineStr == "+++" {
				if readingMetadata {
					readingMetadata = false
				} else {
					readingMetadata = true
				}

				continue
			}

			if readingMetadata {
				iEqual := strings.Index(lineStr, "=")

				if iEqual == -1 {
					panicInvalidLine(lineStr, inputArticleEntry.Name())
				}

				tag := lineStr[:iEqual]
				value := lineStr[iEqual+1:]

				switch tag {
				case "TITLE":
					outputArticle.WriteString(createHeader(value))
				case "DESCRIPTION":
					articleData.Description = value
				case "DATE":
					articleData.Date = value
				default:
					panicInvalidLine(lineStr, inputArticleEntry.Name())
				}
			} else {
				if iStart := strings.Index(lineStr, "{{"); iStart != -1 {
					iEnd := strings.Index(lineStr, "}}")
					iEqual := strings.Index(lineStr, "=")

					if iEnd == -1 || iEqual == -1 {
						panicInvalidLine(lineStr, outputArticle.Name())
					}

					outputArticle.WriteString(lineStr[:iStart])

					tag := lineStr[iStart+2 : iEqual]
					value := lineStr[iEqual+1 : iEnd]

					switch tag {
					case "IMAGE":
						panic("not implemented")
					case "LINK":
						outputArticle.WriteString(createLink(value))
					case "FILE":
						panic("not implemented")
					}

					outputArticle.WriteString(lineStr[iEnd+2:])
				} else {
					if lineStr == "" {
						outputArticle.WriteString(breakLine)
						outputArticle.WriteString(breakLine)
					}

					outputArticle.WriteString(lineStr)
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
