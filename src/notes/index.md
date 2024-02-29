---
layout: layout.html
eleventyExcludeFromCollections: true
permalink: "notes/index.html"
---

# Notes.

{% for note in collections.notes %}

- [#{{ forloop.index }}:]({{ note.url }}) {{ note.data.title }}

{% endfor %}
