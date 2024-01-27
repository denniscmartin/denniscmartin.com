package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"regexp"
	"strings"
)

const (
	METADATA_TAG_END   = "+++"
	METADATA_TAG_TITLE = "TITLE"
	METADATA_TAG_DESCR = "DESCRIPTION"
	METADATA_TAG_DATE  = "DATE"
	TEMPLATE_INDEX     = "../web/src/templates/index.html"
	TEMPLATE_ARTICLE   = "../web/src/templates/article.html"
	IN_ARTICLES_DIR    = "../web/src/articles"
	IN_STATIC_DIR      = "../web/src/static"
	OUT_DIR            = "../web/out"
	OUT_INDEX          = "../web/out/index.html"
	OUT_ARTICLES_DIR   = "../web/out/articles"
	OUT_STATIC_DIR     = "../web/out/static"
)

type Article struct {
	Title       string
	Description string
	Date        string
	Url         string
	HtmlContent string
}

type Paragraph struct {
	text        string
	isCodeBlock bool
}

type PatternsTable struct {
	patterns map[string]Pattern
}

type Pattern struct {
	regex         string
	convertToHtml func(matches []string) string
}

var patternsTable = PatternsTable{
	patterns: map[string]Pattern{
		"MD_HEADING": {
			regex: `(?m)^(#{1,6}) (.+)$`,
			convertToHtml: func(matches []string) string {
				headingCount := len(matches[1])
				return fmt.Sprintf(`<h%d>%s</h%d>`, headingCount, matches[2], headingCount)
			},
		},
		"MD_IMAGE": {
			regex: `(?m)!\[(.*?)\]\((.*?)\)`,
			convertToHtml: func(matches []string) string {
				return fmt.Sprintf(`<img src="%s" alt="%s" with="300">`, matches[2], matches[1])
			},
		},
		"MD_LINK": {
			regex: `\[(.+?)\]\((.+?)\)`,
			convertToHtml: func(matches []string) string {
				return fmt.Sprintf(`<a href="%s" target="_blank">%s</a>`, matches[2], matches[1])
			},
		},
		"MD_LINK_SIMPLE": {
			regex: `(?m)<([^>]+)>`,
			convertToHtml: func(matches []string) string {
				return fmt.Sprintf(`<a href="%s" target="_blank">%s</a>`, matches[1], matches[1])
			},
		},
		"MD_CODE_BLOCK": {
			regex: "```([^`]+?)```",
			convertToHtml: func(matches []string) string {
				var html string
				scanner := bufio.NewScanner(strings.NewReader(matches[1]))

				for scanner.Scan() {
					line := scanner.Text()

					var counter int
					var countingIndentation bool

					for i, r := range line {
						if i == 0 {
							countingIndentation = true
						}

						if countingIndentation {
							if r == ' ' {
								counter++
							} else {
								countingIndentation = false
								break
							}
						}
					}

					line = strings.TrimSpace(line)

					if counter == 0 {
						html += line
					} else {
						html += fmt.Sprintf(`<span style="margin-left: %dpx">%s</span>`, counter*10, line)
					}

					html += `<br>`
				}

				return fmt.Sprintf(`<div class="code-block">%s</div>`, html)
			},
		},
		"MD_CODE_INLINE": {
			regex: "`([^`]+?)`",
			convertToHtml: func(matches []string) string {
				return fmt.Sprintf(`<span class="code-inline">%s</span>`, matches[1])
			},
		},
	},
}

func panicInvalidLine(line string, filename string) {
	panic(fmt.Sprintf("invalid line: %s in file %s\n", line, filename))
}

func splitParagraphs(text string) []Paragraph {
	var paragraphs []Paragraph

	scanner := bufio.NewScanner(strings.NewReader(text))

	var inCodeBlock bool
	var paragraph Paragraph

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "```") {
			inCodeBlock = true
		}

		if inCodeBlock {
			paragraph.text += line
			paragraph.text += "\r\n"

			if strings.HasPrefix(line, "```") {
				inCodeBlock = false
				paragraph.isCodeBlock = true
				paragraphs = append(paragraphs, paragraph)
				paragraph = Paragraph{}
			}
		} else {
			if line == "" {
				paragraphs = append(paragraphs, paragraph)
				paragraph = Paragraph{}
			} else {
				paragraph.text += line
				paragraph.text += "\r\n"
			}
		}
	}

	paragraphs = append(paragraphs, paragraph)

	return paragraphs
}

func main() {
	deleteDir(OUT_DIR)
	createDir(OUT_DIR)
	createDir(OUT_ARTICLES_DIR)
	createDir(OUT_STATIC_DIR)

	var inArticlesDir []fs.DirEntry

	for _, file := range readDir(IN_ARTICLES_DIR) {
		if file.IsDir() {
			continue
		}

		if strings.HasSuffix(file.Name(), ".md") {
			inArticlesDir = append(inArticlesDir, file)
		}
	}

	var articles []Article

	for _, inArticleEntry := range inArticlesDir {
		var article Article
		inArticle := string(readFile(fmt.Sprintf("%s/%s", IN_ARTICLES_DIR, inArticleEntry.Name())))
		inParticleParts := strings.Split(inArticle, METADATA_TAG_END)
		metadata := inParticleParts[0]
		metadataScanner := bufio.NewScanner(strings.NewReader(metadata))

		for metadataScanner.Scan() {
			line := metadataScanner.Text()
			iEqual := strings.Index(line, "=")

			if iEqual == -1 {
				panicInvalidLine(line, inArticleEntry.Name())
			}

			tag := line[:iEqual]
			value := line[iEqual+1:]

			switch tag {
			case METADATA_TAG_TITLE:
				article.Title = value
			case METADATA_TAG_DESCR:
				article.Description = value
			case METADATA_TAG_DATE:
				article.Date = value
			default:
				panicInvalidLine(line, inArticleEntry.Name())
			}
		}

		body := strings.TrimSpace(inParticleParts[1])
		paragraphs := splitParagraphs(body)
		var htmlBody string

		for _, paragraph := range paragraphs {
			var submatchesCounter int

			htmlParagraph := paragraph.text

			for patternName, pattern := range patternsTable.patterns {
				if paragraph.isCodeBlock && patternName != "MD_CODE_BLOCK" {
					continue
				}

				reg := regexp.MustCompile(pattern.regex)
				submatches := reg.FindAllStringSubmatch(paragraph.text, -1)
				submatchesCounter += len(submatches)

				for _, matches := range submatches {
					htmlTag := pattern.convertToHtml(matches)
					htmlParagraph = strings.ReplaceAll(htmlParagraph, matches[0], htmlTag)
					paragraph.text = strings.ReplaceAll(paragraph.text, matches[0], "")
				}
			}

			if strings.HasPrefix(htmlParagraph, "<h") || strings.HasPrefix(htmlParagraph, "<div") {
				htmlBody += htmlParagraph
			} else {
				htmlBody += `<p>`
				htmlBody += htmlParagraph
				htmlBody += `</p>`
			}
		}

		article.HtmlContent = string(readFile(TEMPLATE_ARTICLE))
		article.HtmlContent = strings.ReplaceAll(article.HtmlContent, "{{TITLE}}", article.Title)
		article.HtmlContent = strings.ReplaceAll(article.HtmlContent, "{{ARTICLE}}", htmlBody)
		outArticleFilename := strings.ReplaceAll(inArticleEntry.Name(), ".md", ".html")
		article.Url = fmt.Sprintf("articles/%s", outArticleFilename)

		writeFile([]byte(article.HtmlContent), fmt.Sprintf("%s/%s", OUT_ARTICLES_DIR, outArticleFilename))

		articles = append(articles, article)
	}

	var htmlArticlesList string

	for _, article := range articles {
		htmlArticlesList += fmt.Sprintf(`<li>%s: <a href="%s">%s</a></li>`,
			article.Date, article.Url, article.Description,
		)
	}

	htmlIndexTemplate := string(readFile(TEMPLATE_INDEX))
	htmlIndexTemplate = strings.ReplaceAll(htmlIndexTemplate, "{{ARTICLES}}", htmlArticlesList)
	writeFile([]byte(htmlIndexTemplate), OUT_INDEX)

	for _, staticItem := range readDir(IN_STATIC_DIR) {
		if staticItem.IsDir() {
			staticDir := readDir(fmt.Sprintf("%s/%s", IN_STATIC_DIR, staticItem.Name()))

			for _, nestedStaticItem := range staticDir {
				file := readFile(fmt.Sprintf("%s/%s/%s",
					IN_STATIC_DIR, staticItem.Name(), nestedStaticItem.Name()))

				writeFile(file, fmt.Sprintf("%s/%s", OUT_STATIC_DIR, nestedStaticItem.Name()))
			}
		} else {
			file := readFile(fmt.Sprintf("%s/%s", IN_STATIC_DIR, staticItem.Name()))
			writeFile(file, fmt.Sprintf("%s/%s", OUT_STATIC_DIR, staticItem.Name()))
		}
	}
}
