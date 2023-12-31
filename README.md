# thetucks.com static site (updated Dec 2023)

[![Deploy Hugo site to Pages](https://github.com/tuck1s/quickstart/actions/workflows/hugo.yaml/badge.svg)](https://github.com/tuck1s/quickstart/actions/workflows/hugo.yaml)

## Structure

This enables standard [Ananke theme](https://github.com/theNewDynamic/gohugo-theme-ananke) to be used.
```
content
  post
    postname
      index.md
      img
        img1.jpg
        : etc
```

## Local development

In project folder, run `hugo server` and open locally in browser.

## Production build

Build on checkin is automatically done using a standard Github Action for Hugo sites.
