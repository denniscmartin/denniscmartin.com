---
layout: layout.html
eleventyExcludeFromCollections: true
permalink: "posts/index.html"
---

# Posts.

{% for post in collections.posts %}

- [{{ post.date | date: "%Y-%m-%d" }}:]({{ post.url }}) {{ post.data.title }}

{% endfor %}
