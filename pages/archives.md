---
layout: page
title: Archives
permalink: /archives/
---

<section id="archive">
    {%for post in site.posts %}
    {% unless post.next %}
    <h4>{{ post.date | date: '%Y' }}</h4>
    <ul>
        {% else %}
        {% capture year %}{{ post.date | date: '%Y' }}{% endcapture %}
        {% capture nyear %}{{ post.next.date | date: '%Y' }}{% endcapture %}
        {% if year != nyear %}
    </ul>
    <h4>{{ post.date | date: '%Y' }}</h4>
    <ul>
        {% endif %}
        {% endunless %}
        <li>
            <time>{{ post.date | date:"%d %b" }}</time>
            <a href="{{ post.url }}">{{ post.title }}</a>
        </li>
        {% endfor %}
    </ul>
</section>

---

{% include post_internal_links.html %}
