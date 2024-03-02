+++
title = '{{ replace .File.ContentBaseName "-" " " | title }}'
date = {{ .Date }}
thumbnail = '{{ printf "thumbnail-%s.png" .File.ContentBaseName }}'
draft = true
+++
