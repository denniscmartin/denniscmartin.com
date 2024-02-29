---
layout: layout.html
pagination:
    data: collections
    size: 1
    alias: tag
    filter:
        - notes
        - posts
        - projects
        - all
permalink: /tags/{{ tag }}/
---

# #{{ tag }}

## Posts

{% for item in collections[ tag ] %}

{% if item.data.tags[0] == 'posts' %}

-   [{{ item.data.published }}:]({{ item.url }}) {{ item.data.title }}

{% endif %}

{% endfor %}

## Notes

{% for item in collections[ tag ] %}

{% if item.data.tags[0] == 'notes' %}

-   [#{{ item.data.number }}:]({{ item.url }}) {{ item.data.title }}

{% endif %}

{% endfor %}

