---
layout: post
title: Create Github Pages with Jekyll
date: 2023-03-17
tags: jekyll
---

Create Github Pages with Jekyll. 😆 Hi....

[About GitHub Pages and Jekyll](https://help.github.com/articles/adding-jekyll-plugins-to-a-github-pages-site/){:target="_blank"}

Jekyll is a static site generator with built-in support for GitHub Pages and a simplified build process. Jekyll takes Markdown and HTML files and creates a complete static website based on your choice of layouts. Jekyll supports Markdown and Liquid, a templating language that loads dynamic content on your site. For more information, see Jekyll.

Jekyll is not officially supported for Windows. For more information, see "Jekyll on Windows" in the Jekyll documentation.

We recommend using Jekyll with GitHub Pages. If you prefer, you can use other static site generators or customize your own build process locally or on another server. For more information, see "About GitHub Pages."


Some configuration settings cannot be changed for GitHub Pages sites.

```bash
  lsi: false
  safe: true
  source: [your repo's top level directory]
  incremental: false
  highlighter: rouge
  gist:
    noscript: false
  kramdown:
    math_engine: mathjax
    syntax_highlighter: rouge
```

GitHub Pages uses plugins that are enabled by default and cannot be disabled:

- jekyll-coffeescript
- jekyll-default-layout
- jekyll-gist
- jekyll-github-metadata
- jekyll-optional-front-matter
- jekyll-paginate
- jekyll-readme-index
- jekyll-titles-from-headings
- jekyll-relative-links

Building your site locally

If you are publishing from a branch, changes to your site are published automatically when the changes are merged into your site's publishing source. If you are publishing from a custom GitHub Actions workflow, changes are published whenever your workflow is triggered (typically by a push to the default branch). If you want to preview your changes first, you can make the changes locally instead of on GitHub. Then, test your site locally.