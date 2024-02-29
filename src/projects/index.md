---
layout: layout.html
eleventyExcludeFromCollections: true
---

# Projects.

{% for project in collections.projects %}

- [{{ project.data.title }}]({{ project.url }})

{% endfor %}
