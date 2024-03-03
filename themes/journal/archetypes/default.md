+++
title = '{{ replace .File.ContentBaseName "-" " " | title }}'
date = {{ .Date }}
thumbnail = '{{ replace .File.Dir "/" "-"  }}{{ .File.BaseFileName }}-thumbnail.png'
draft = false
+++
