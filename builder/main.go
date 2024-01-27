package main

import (
	"bufio"
	"fmt"
	"regexp"
	"sort"
	"strings"
)

const (
	METADATA_TAG_END   = "+++"
	METADATA_TAG_TITLE = "TITLE"
	METADATA_TAG_DESCR = "DESCRIPTION"
	METADATA_TAG_DATE  = "DATE"
	TEMPLATE_INDEX     = "../web/templates/index.html"
	TEMPLATE_ARTICLE   = "../web/templates/article.html"
	IN_ARTICLES_DIR    = "../web/articles"
	OUT_INDEX          = "../web/out/index.html"
	OUT_ARTICLES_DIR   = "../web/out/articles"
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

func splitParagraphs(text string) []string {
	var paragraphs []string

	scanner := bufio.NewScanner(strings.NewReader(text))

	var inCodeBlock bool
	var paragraph string

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "```") {
			inCodeBlock = !inCodeBlock
			paragraph += line
			paragraph += "\r\n"
			continue
		}

		if line == "" && !inCodeBlock {
			paragraphs = append(paragraphs, paragraph)
			paragraph = ""
			continue
		}

		paragraph += line
		// FIXME: Does it works on linux (\r\n\r\n)?
		paragraph += "\r\n"
	}

	paragraphs = append(paragraphs, paragraph)

	return paragraphs
}

func main() {
	inArticlesDir := readDir(IN_ARTICLES_DIR)

	sort.Slice(inArticlesDir, func(i, j int) bool {
		return inArticlesDir[i].Name() > inArticlesDir[j].Name()
	})

	var articles []Article

	for _, inArticleEntry := range inArticlesDir {
		var article Article
		inArticle := readFile(fmt.Sprintf("%s/%s", IN_ARTICLES_DIR, inArticleEntry.Name()))
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

			htmlParagraph := paragraph
			for _, pattern := range patternsTable.patterns {
				reg := regexp.MustCompile(pattern.regex)
				submatches := reg.FindAllStringSubmatch(paragraph, -1)
				submatchesCounter += len(submatches)

				for _, matches := range submatches {
					htmlTag := pattern.convertToHtml(matches)
					htmlParagraph = strings.ReplaceAll(htmlParagraph, matches[0], htmlTag)
					paragraph = strings.ReplaceAll(paragraph, matches[0], "")
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

		article.HtmlContent = readFile(TEMPLATE_ARTICLE)
		article.HtmlContent = strings.ReplaceAll(article.HtmlContent, "{{TITLE}}", article.Title)
		article.HtmlContent = strings.ReplaceAll(article.HtmlContent, "{{ARTICLE}}", htmlBody)
		outArticleFilename := strings.ReplaceAll(inArticleEntry.Name(), ".md", ".html")
		article.Url = fmt.Sprintf("articles/%s", outArticleFilename)

		writeFile(article.HtmlContent, fmt.Sprintf("%s/%s", OUT_ARTICLES_DIR, outArticleFilename))

		articles = append(articles, article)
	}

	var htmlArticlesList string

	for _, article := range articles {
		htmlArticlesList += fmt.Sprintf(`<li>%s: <a href="%s">%s</a></li>`,
			article.Date, article.Url, article.Description,
		)
	}

	htmlIndexTemplate := readFile(TEMPLATE_INDEX)
	htmlIndexTemplate = strings.ReplaceAll(htmlIndexTemplate, "{{ARTICLES}}", htmlArticlesList)
	writeFile(htmlIndexTemplate, OUT_INDEX)
}
