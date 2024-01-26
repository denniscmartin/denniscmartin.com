package main

import "fmt"

const openHtml = `
<!doctype html>
<html>
`

const closeHtml = `
</html>
`

const head = `
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
`

const openBody = `
<body>
`

const closeBody = `
</body>
`

const navbar = `
<nav>
	<li><a href="/">Go to home</a>
</nav>
`

const openMain = `
<main>
`

const closeMain = `
</main>
`

const breakLine = `
<br>
`

func createHeader(title string) string {
	return fmt.Sprintf(`
<div class="header">
	<h1 class="title bordered">%s</h1>
</div>
`, title)
}

func createLink(url string) string {
	return fmt.Sprintf(`
<a href=%s target="_blank">%s</a>
`, url, url)
}

func createArticleListItem(article ArticleItem) string {
	return fmt.Sprintf(`
<li>%s: <a href="%s">%s</a></li>	
`, article.Date, article.Url, article.Description)
}
