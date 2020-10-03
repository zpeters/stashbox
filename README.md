# stashbox

[![Go Report Card](https://goreportcard.com/badge/github.com/zpeters/stashbox)](https://goreportcard.com/report/github.com/zpeters/stashbox)
[![Build Status](https://travis-ci.org/zpeters/stashbox.svg?branch=main)](https://travis-ci.org/zpeters/stashbox)
![Run Gosec](https://github.com/zpeters/stashbox/workflows/Run%20Gosec/badge.svg?branch=main)
![CodeQL](https://github.com/zpeters/stashbox/workflows/CodeQL/badge.svg)
[![CII Best Practices](https://bestpractices.coreinfrastructure.org/projects/4318/badge)](https://bestpractices.coreinfrastructure.org/projects/4318)
[![License](https://img.shields.io/github/license/zpeters/stashbox)](https://img.shields.io/github/license/zpeters/stashbox)
[![Contributers](https://img.shields.io/github/contributors/zpeters/stashbox)](https://img.shields.io/github/contributors/zpeters/stashbox)

## Stashbox is your personal Internet Archive

The goal of stashbox is to help you create your own - personal copy of websites that you wish to archive.

The initial way to do this will be to run a simple command, but in the future it can be extended to a web interface, a plugin or other means.

Having a local "static" copy of a website can help for research, change tracking and for many other purposes

## Roadmap

- [x] Initial command line client to add urls to a personal archive with Text, Html and Pdf copies of the website
- [x] Ability to save new versions of the same URL
- [ ] Version "diffing" and browsing
- [ ] User friendly interface (web, etc)
- [ ] Searching and other functions

## Usage flags

```
Usage of ./stashbox:
  -b string
    	folder to save stash into (default "./stashDb")
  -crawl
    	crawl and save websites
  -list
    	list saved archives
```

## Instructions for usage

1. Build the binary with name preferably stashbox
1. You can give path to the stash folder using `-b` flag else it will default to ./stashDb
1. Do `./stashbox -list` to get list of saved websites
1. To save a single or multiple websites you have 2 options to provide the inputs

- Run `./stashbox -crawl` and then provide number of websites and the urls of the websites to save as command line inputs
- Create a urls.txt file with following structure and feed it to the binary like `./stashbox -crawl < urls.txt`
  ```
  5
  https://www.site1.com
  https://www.site2.com
  https://www.site3.com
  https://www.site4.com
  https://www.site5.com
  ```

## Contributing

New issues and pull requests are welcomed. Please see [CONTRIBUTING.md](CONTRIBUTING.md)
