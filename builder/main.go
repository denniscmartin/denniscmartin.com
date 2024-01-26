package main

import (
	"bufio"
	"fmt"
	"regexp"
	"sort"
	"strings"
)

const (
	METADATA_END          = "+++"
	METADATA_TITLE        = "TITLE"
	METADATA_DESCR        = "DESCRIPTION"
	METADATA_DATE         = "DATE"
	HTML_ARTICLE_TEMPLATE = `
	<!doctype html>
	<html>
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<link href="../styles.css" rel="stylesheet">
			<link rel="preconnect" href="https://fonts.googleapis.com">
			<link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
			<link
				href="https://fonts.googleapis.com/css2?family=IBM+Plex+Sans:ital,wght@0,100;0,200;0,300;0,400;0,500;0,600;0,700;1,100;1,200;1,300;1,400;1,500;1,600;1,700&display=swap"
				rel="stylesheet">
		</head>
		<body>
			<div class="header">
				<h1 class="title bordered">{{TITLE}}</h1>
			</div>
			<nav>
				<li><a href="/">Go to home</a></li>
			</nav>
			<main>
				{{ARTICLE}}
			</main>
		</body>
	</html>
	`
)

type Article struct {
	Title       string
	Description string
	Date        string
	Url         string
	HtmlContent string
}

type PatternsTable struct {
	patterns map[string]Pattern
}

type Pattern struct {
	regex   string
	convert func(matches []string) string
}

var patternsTable = PatternsTable{
	patterns: map[string]Pattern{
		"MD_LINK": {
			regex: `\[(.+?)\]\((.+?)\)`,
			convert: func(matches []string) string {
				return fmt.Sprintf(`<a href="%s" target="_blank">%s</a>`, matches[2], matches[1])
			},
		},
	},
}

func panicInvalidLine(line string, filename string) {
	panic(fmt.Sprintf("invalid line: %s in file %s\n", line, filename))
}

func main() {
	inputArticlesDir := readInArticlesDir()

	sort.Slice(inputArticlesDir, func(i, j int) bool {
		return inputArticlesDir[i].Name() > inputArticlesDir[j].Name()
	})

	var articles []Article

	for _, inputArticleEntry := range inputArticlesDir {
		var article Article
		inputArticle := strings.Split(readInArticle(inputArticleEntry.Name()), METADATA_END)
		metadata := inputArticle[0]
		metadataScanner := bufio.NewScanner(strings.NewReader(metadata))

		for metadataScanner.Scan() {
			line := metadataScanner.Text()
			iEqual := strings.Index(line, "=")

			if iEqual == -1 {
				panicInvalidLine(line, inputArticleEntry.Name())
			}

			tag := line[:iEqual]
			value := line[iEqual+1:]

			switch tag {
			case METADATA_TITLE:
				article.Title = value
			case METADATA_DESCR:
				article.Description = value
			case METADATA_DATE:
				article.Date = value
			default:
				panicInvalidLine(line, inputArticleEntry.Name())
			}
		}

		mdArticle := inputArticle[1]

		for _, regexPattern := range patternsTable.patterns {
			reg := regexp.MustCompile(regexPattern.regex)
			submatches := reg.FindAllStringSubmatch(mdArticle, -1)

			if len(submatches) > 0 {
				for _, matches := range submatches {
					newStr := regexPattern.convert(matches)
					mdArticle = strings.ReplaceAll(mdArticle, matches[0], newStr)
				}
			}
		}

		mdArticle = strings.TrimSpace(mdArticle)

		bodyScanner := bufio.NewScanner(strings.NewReader(mdArticle))
		var lineCounter int
		var htmlArticle string

		for bodyScanner.Scan() {
			line := bodyScanner.Text()

			if lineCounter == 0 {
				htmlArticle += `<p>`
			}

			if line == "" {
				htmlArticle += `</p><p>`
			}

			htmlArticle += line
			lineCounter++
		}

		htmlArticle += `</p>`

		article.HtmlContent = strings.ReplaceAll(HTML_ARTICLE_TEMPLATE, "{{TITLE}}", article.Title)
		article.HtmlContent = strings.ReplaceAll(article.HtmlContent, "{{ARTICLE}}", htmlArticle)
		outArticle, url := createOutArticle(inputArticleEntry.Name())
		article.Url = url
		outArticle.WriteString(article.HtmlContent)

		articles = append(articles, article)
	}

	var htmlArticlesList string

	for _, article := range articles {
		htmlArticlesList += fmt.Sprintf(`<li>%s: <a href="%s">%s</a></li>`,
			article.Date, article.Url, article.Description,
		)
	}

	htmlIndexTemplate := readInIndexFile()
	htmlIndexTemplate = strings.ReplaceAll(htmlIndexTemplate, "{{ARTICLES}}", htmlArticlesList)
	writeOutIndexFile(htmlIndexTemplate)
}
