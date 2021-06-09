---
layout: page
title: Jay's Space
subtitle: Just give myself another reason to move on.
---

<ul>
  {% for post in site.posts %}
    <li>
      <a href="{{ post.url }}">{{ post.title }}</a>
    </li>
  {% endfor %}
</ul>