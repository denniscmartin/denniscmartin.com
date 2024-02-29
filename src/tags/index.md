---
layout: layout.html
---

# Tags

{% for tag in collections %}
{% if tag[0] != 'notes' and tag[0] != 'projects' and tag[0] != 'all' and tag[0] != 'posts' %}
- [#{{ tag[0] }}](/tags/{{ tag[0] }})
{% endif %}
{% endfor %}