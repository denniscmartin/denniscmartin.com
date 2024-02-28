---
layout: layout.html
eleventyExcludeFromCollections: true
---

# Posts.

{% for post in collections.posts %}

- {{ post.data.published }}: {{ post.data.title }}

{% endfor %}
