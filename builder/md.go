package main

const (
	METADATA_END   = "+++"
	METADATA_TITLE = "TITLE"
	METADATA_DESCR = "DESCRIPTION"
	METADATA_DATE  = "DATE"

	MD_HEADING    = "#"
	MD_BOLD       = "**"
	MD_ITALIC     = "*"
	MD_BLOCKQUOTE = ">"
	MD_CODE       = "`"
	MD_CODE_BLOCK = "```"
	MD_LINK_OPEN  = "<"
	MD_LINK_CLOSE = ">"
	MD_IMAGE      = "!"
)

type converter struct {
	md map[string]string
}

var conv converter = converter{
	
}