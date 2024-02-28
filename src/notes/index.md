---
layout: layout.html
eleventyExcludeFromCollections: true
---

# Notes.

{% for note in collections.notes %}

- #{{ note.data.number }}: {{ note.data.title }}

{% endfor %}
