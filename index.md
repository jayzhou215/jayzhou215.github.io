---
layout: page
title: Jay's Space
subtitle: 阿杰的个人空间
---

## 前言
记录自己的学习心得，Do not just see see, try it yourself.


## go_study
[detail](./go_study/index.md)

## 算法心得
[algorithm thinks](./algorithm/visualgo.md)

## 架构心得
[architecture thinks](./architecture/index.md)

<ul>
  {% for post in site.posts %}
    <li>
      <a href="{{ post.url }}">{{ post.title }}</a>
    </li>
  {% endfor %}
</ul>