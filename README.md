# stashbox
![Go Build](https://github.com/zpeters/stashbox/workflows/Go%20Build/badge.svg)
![Go Test](https://github.com/zpeters/stashbox/workflows/Go%20Test/badge.svg)
![Go Lint](https://github.com/zpeters/stashbox/workflows/Go%20Lint/badge.svg)
![Gosec](https://github.com/zpeters/stashbox/workflows/Gosec/badge.svg)
![CodeQL](https://github.com/zpeters/stashbox/workflows/CodeQL/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/zpeters/stashbox)](https://goreportcard.com/report/github.com/zpeters/stashbox)
[![License](https://img.shields.io/github/license/zpeters/stashbox)](https://img.shields.io/github/license/zpeters/stashbox)
[![Contributers](https://img.shields.io/github/contributors/zpeters/stashbox)](https://img.shields.io/github/contributors/zpeters/stashbox)
![gopherbadger-tag-do-not-edit](https://github.com/KrishKayc/stashbox/blob/add-code-coverage-badge/coverage/coverage_badge.png)
## Stashbox is your personal Internet Archive

The goal of stashbox is to help you create your own - personal copy of websites that you wish to archive.

The initial way to do this will be to run a simple command, but in the future it can be extended to a web interface, a plugin or other means.

Having a local "static" copy of a website can help for research, change tracking and for many other purposes

## Roadmap

- [x]  Initial command line client to add urls to a personal archive with Text, Html and Pdf copies of the website
- [x]  Ability to save new versions of the same URL
- [ ]  Version "diffing" and browsing
- [ ]  User friendly interface (web, etc)
- [ ]  Searching and other functions

## Usage
```
Usage: stashbox <command> <options>

  Where command is one of:
    add   --  add a url to the archive
      -b string
            stashbox archive directory (defaults to ./stashDb) (default "./stashDb")
      -u string
            url to download
    list  --  list all archives
      -b string
            stashbox archive directory (defaults to ./stashDb) (default "./stashDb")
    open  --  open an archive
      -b string
            stashbox archive directory (defaults to ./stashDb) (default "./stashDb")
      -n int
            archive number to open (from list command)

  To see help text, you can run:
    stashbox <command> -h
```

## Contributing

New issues and pull requests are welcomed.  Please see [CONTRIBUTING.md](CONTRIBUTING.md)

## Contributors

<a href="https://github.com/zpeters/stashbox/graphs/contributors">
  <img src="https://contributors-img.web.app/image?repo=zpeters/stashbox" />
</a>

Made with [contributors-img](https://contributors-img.web.app).
